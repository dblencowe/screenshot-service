package main

import (
	"log"
	"net/http"
	"os"
	"screenshot-service/pkg/endpoints"
)

func main() {
	http.HandleFunc("/screenshot", endpoints.Screenshot)

	log.Printf("http server listening on %s", os.Getenv("HTTP_ADDR"))
	if err := http.ListenAndServe(os.Getenv("HTTP_ADDR"), nil); err != nil {
		log.Fatal(err)
	}
}
