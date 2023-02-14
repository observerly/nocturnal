package query

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

	// Setup the route:
	r.GET("/default", func(c *gin.Context) {
		datetime, longitude, latitude := GetDefaultObserverParams(c)
		t.Log(datetime, longitude, latitude)
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
}

func TestGetDefaultObserverQueryWhenPopulated(t *testing.T) {
	// Set gin mode to test mode:
	gin.SetMode(gin.TestMode)

	// Setup the default router:
	r := gin.Default()

	// Setup the route:
	r.GET("/default", func(c *gin.Context) {
		datetime, longitude, latitude := GetDefaultObserverParams(c)
		t.Log(datetime, longitude, latitude)
	})

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, "/default?datetime=2020-01-01T00:00:00Z&longitude=1&latitude=1", nil)

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
}
