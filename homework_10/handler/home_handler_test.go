package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := &http.Request{}
	rw := httptest.NewRecorder()

	handler := NewHomeHandler(nil)
	handler.ServeHTTP(rw, req)

	expectedBody := "Hello!"

	if rw.Body.String() != expectedBody {
		t.Errorf("invalid response body: got: %s, want: %s", rw.Body, expectedBody)
	}
}
