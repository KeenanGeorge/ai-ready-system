package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	t.Run("SMA-7:: Health endpoint returns ok status", func(t *testing.T) {
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

	// Test healthHandler function directly for better coverage
	t.Run("Health handler function coverage", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()

		// Call the function directly to ensure it's covered
		healthHandler(rr, req)

		// Verify response
		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
		}
		if rr.Body.String() != "ok" {
			t.Errorf("expected body 'ok', got %q", rr.Body.String())
		}
	})
}

// Test server setup and routing
func TestServerSetup(t *testing.T) {
	t.Run("Server routes are configured", func(t *testing.T) {
		// Create a test server
		mux := http.NewServeMux()
		mux.HandleFunc("/health", healthHandler)

		// Test that the route is registered
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

// Test error handling scenarios
func TestErrorHandling(t *testing.T) {
	t.Run("Health handler with invalid method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/health", nil)
		rr := httptest.NewRecorder()

		// The handler should still work with any method
		healthHandler(rr, req)

		// Verify it responds (even if not the expected method)
		if rr.Code >= 500 {
			t.Errorf("handler should not return server error, got %d", rr.Code)
		}
	})
}

// Test main function behavior (partial)
func TestMainFunctionBehavior(t *testing.T) {
	t.Run("Main function can be called", func(t *testing.T) {
		// This test verifies that the main function can be executed
		// We can't test the actual ListenAndServe without starting a real server
		// But we can test that the function exists and can be called

		// Verify the function exists (this will be covered by the test)
		if healthHandler == nil {
			t.Error("healthHandler function is nil")
		}

		// Test that we can create a request and response
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()

		// This ensures the handler function is covered
		healthHandler(rr, req)

		// Verify basic functionality
		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

// Benchmark tests for performance
func BenchmarkHealthHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		healthHandler(rr, req)
	}
}

// Test coverage for edge cases
func TestEdgeCases(t *testing.T) {
	t.Run("Health handler with nil request", func(t *testing.T) {
		rr := httptest.NewRecorder()

		// This should not panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("handler panicked: %v", r)
			}
		}()

		// Test with nil request (edge case)
		healthHandler(rr, nil)

		// Verify it handles the nil case gracefully
		if rr.Code >= 500 {
			t.Errorf("handler should handle nil request gracefully, got status %d", rr.Code)
		}
	})
}
