package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()
	healthHandler(rr, req)
	if rr.Body.String() != "ok" {
		t.Fatalf("expected ok, got %q", rr.Body.String())
	}
}
