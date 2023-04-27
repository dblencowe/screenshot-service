package endpoints_test

import (
	"net/http"
	"net/http/httptest"
	"screenshot-service/pkg/endpoints"
	"testing"
)

func TestJSONError(t *testing.T) {
	req, err := http.NewRequest("POST", "/screenshot", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		endpoints.JSONError(w, "an expected error", http.StatusBadRequest)
	})
	handler.ServeHTTP(rr, req)

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json; charset=utf-8" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json; charset=utf-8")
	}

	if rr.Code != http.StatusBadRequest {
		t.Errorf("status code does not match: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
}
