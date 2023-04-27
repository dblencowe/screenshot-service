package endpoints

import (
	"encoding/json"
	"fmt"
	"mvdan.cc/xurls/v2"
	"net/http"
	"screenshot-service/pkg/screenshot"
)

type ScreenshotRequestBody struct {
	Url string `json:"url"`
}

func Screenshot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if req.Method != http.MethodPost {
		JSONError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var r ScreenshotRequestBody
	if req.Body == nil {
		JSONError(w, "missing json body", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	rxStrict := xurls.Strict()
	url := rxStrict.FindString(r.Url)
	if len(url) == 0 {
		JSONError(w, fmt.Sprintf("unable to find a valid url in %s", r.Url), http.StatusBadRequest)
		return
	}

	result, err := screenshot.MakeScreenshot(url)
	if err != nil {
		JSONError(w, fmt.Sprintf("error creating screenshot: %s", err), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONError(w, fmt.Sprintf("error encoding response: %s", err), http.StatusInternalServerError)
	}
}
