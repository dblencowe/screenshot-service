package endpoints_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"screenshot-service/pkg/endpoints"
	"testing"
)

func TestScreenshotHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		body := bytes.NewReader([]byte(`{"url":"https://github.com/arborealfirecat"}`))
		req, err := http.NewRequest("POST", "/screenshot", body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(endpoints.Screenshot)
		handler.ServeHTTP(rr, req)

		if ctype := rr.Header().Get("Content-Type"); ctype != "application/json; charset=utf-8" {
			t.Errorf("content type header does not match: got %v want %v", ctype, "application/json; charset=utf-8")
		}

		if rr.Code != http.StatusOK {
			t.Errorf("status code does not match: got %v, want %v", rr.Code, http.StatusOK)
		}
	})

	t.Run("no body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/screenshot", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(endpoints.Screenshot)
		handler.ServeHTTP(rr, req)

		if ctype := rr.Header().Get("Content-Type"); ctype != "application/json; charset=utf-8" {
			t.Errorf("content type header does not match: got %v want %v", ctype, "application/json; charset=utf-8")
		}

		if rr.Code != http.StatusBadRequest {
			t.Errorf("status code does not match: got %v, want %v", rr.Code, http.StatusBadRequest)
		}
	})

	t.Run("wrong method", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/screenshot", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(endpoints.Screenshot)
		handler.ServeHTTP(rr, req)

		if ctype := rr.Header().Get("Content-Type"); ctype != "application/json; charset=utf-8" {
			t.Errorf("content type header does not match: got %v want %v", ctype, "application/json; charset=utf-8")
		}

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("status code does not match: got %v, want %v", rr.Code, http.StatusMethodNotAllowed)
		}
	})
}
