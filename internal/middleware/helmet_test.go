package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetDefaultObserverQueryUndefined(t *testing.T) {
	// Set gin mode to test mode:
	gin.SetMode(gin.TestMode)

	// Setup the default router:
	r := gin.Default()

	// Setup Helmet Security Headers:
	r.Use(HelmetMiddleware())

	// Setup the route:
	r.GET("/default", func(c *gin.Context) {
		t.Log("Default Helmet Middleware Test")
	})

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, "/default", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	// Check to see if the response headers were what you expected
	if w.Header().Get("Content-Security-Policy") != "default-src 'self'" {
		t.Fatalf("Expected to get Content-Security-Policy header %s but instead got %s\n", "default-src 'self'", w.Header().Get("Content-Security-Policy"))
	}

	if w.Header().Get("Cross-Origin-Opener-Policy") != "same-origin" {
		t.Fatalf("Expected to get Cross-Origin-Opener-Policy header %s but instead got %s\n", "same-origin", w.Header().Get("Cross-Origin-Opener-Policy"))
	}

	if w.Header().Get("Referrer-Policy") != "strict-origin-when-cross-origin" {
		t.Fatalf("Expected to get Referrer-Policy header %s but instead got %s\n", "strict-origin-when-cross-origin", w.Header().Get("Referrer-Policy"))
	}

	if w.Header().Get("Strict-Transport-Security") != "max-age=5184000; includeSubDomains" {
		t.Fatalf("Expected to get Strict-Transport-Security header %s but instead got %s\n", "max-age=5184000; includeSubDomains", w.Header().Get("Strict-Transport-Security"))
	}

	if w.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatalf("Expected to get X-Content-Type-Options header %s but instead got %s\n", "nosniff", w.Header().Get("X-Content-Type-Options"))
	}

	if w.Header().Get("X-Download-Options") != "noopen" {
		t.Fatalf("Expected to get X-Download-Options header %s but instead got %s\n", "noopen", w.Header().Get("X-Download-Options"))
	}

	if w.Header().Get("X-DNS-Prefetch-Control") != "off" {
		t.Fatalf("Expected to get X-DNS-Prefetch-Control header %s but instead got %s\n", "off", w.Header().Get("X-DNS-Prefetch-Control"))
	}

	if w.Header().Get("X-Frame-Options") != "Deny" {
		t.Fatalf("Expected to get X-Frame-Options header %s but instead got %s\n", "Deny", w.Header().Get("X-Frame-Options"))
	}

	if w.Header().Get("X-XSS-Protection") != "1; mode=block" {
		t.Fatalf("Expected to get X-XSS-Protection header %s but instead got %s\n", "1; mode=block", w.Header().Get("X-XSS-Protection"))
	}
}
