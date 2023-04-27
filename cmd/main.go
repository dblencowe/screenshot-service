package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mvdan.cc/xurls/v2"
	"net/http"
	"os"
	"screenshot-service/pkg/screenshot"
)

type ScreenshotRequestBody struct {
	Url string `json:"url"`
}

func main() {
	http.HandleFunc("/screenshot", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var r ScreenshotRequestBody
		if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rxStrict := xurls.Strict()
		url := rxStrict.FindString(r.Url)
		if len(url) == 0 {
			http.Error(w, fmt.Sprintf("unable to find a valid url in %s", r.Url), http.StatusBadRequest)
			return
		}

		result, err := screenshot.MakeScreenshot(url)
		if err != nil {
			http.Error(w, fmt.Sprintf("error creating screenshot: %s", err), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, fmt.Sprintf("error encoding response: %s", err), http.StatusInternalServerError)
		}
	})

	log.Printf("http server listening on %s", os.Getenv("HTTP_ADDR"))
	if err := http.ListenAndServe(os.Getenv("HTTP_ADDR"), nil); err != nil {
		log.Fatal(err)
	}
}
