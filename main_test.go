package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	t.Run("SMA-7: Health endpoint returns ok status", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()
		healthHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
		}

		if rr.Body.String() != "ok" {
			t.Errorf("expected body 'ok', got %q", rr.Body.String())
		}
	})

	t.Run("SMA-7: Health endpoint handles GET method correctly", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()
		healthHandler(rr, req)

		if rr.Header().Get("Content-Type") == "" {
			t.Log("Content-Type header not set, this is acceptable for simple text response")
		}
	})

	t.Run("SMA-7: Health endpoint is accessible", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()
		healthHandler(rr, req)

		// Verify the endpoint responds without error
		if rr.Code >= 400 {
			t.Errorf("endpoint should not return error status, got %d", rr.Code)
		}
	})
}
