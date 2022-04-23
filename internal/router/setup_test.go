package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup the Gin API router:
var r = SetupRouter()

// Setup the base response struct:
var response map[string]string

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestSetupRouter(t *testing.T) {
	routes := r.Routes()
	fmt.Println(routes)
}

func TestNoRoute(t *testing.T) {
	// Perform a GET request with that handler.
	w := performRequest(r, "GET", "/")
	// Assert we redirected correctly, the request gives a 302:
	assert.Equal(t, http.StatusFound, w.Code)
}

func TestVersionRoute(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"latest": "/api/v1",
	}

	// Perform a GET request with that handler.
	w := performRequest(r, "GET", "/version")

	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the value & whether or not it exists
	value, exists := response["latest"]

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["latest"], value)
}

func TestAPIBaseRoute(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"description": "Nocturnal 🌑 is observerly's Gin Gonic API for Lunar and Solar advanced scheduling, that utilises Dusk.",
		"endpoint":    "/api/v1",
		"name":        "Nocturnal API by observerly",
	}

	// Perform a GET request with that handler.
	w := performRequest(r, "GET", "/api/v1")

	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the description & whether or not it exists
	description, exists := response["description"]
	assert.True(t, exists)

	// Grab the endpoint & whether or not it exists
	endpoint, exists := response["endpoint"]
	assert.True(t, exists)

	// Grab the name & whether or not it exists
	name, exists := response["name"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, body["description"], description)
	assert.Equal(t, body["endpoint"], endpoint)
	assert.Equal(t, body["name"], name)
}
