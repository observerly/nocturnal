package twilight

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTwilightRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	r.GET("/api/v1/twilight", GetTwilight)

	return r
}

// Setup the Gin API router:
var r = SetupTwilightRouter()

// Setup the base response struct:
var response map[string]map[string]interface{}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var w = performRequest(r, "GET", "/api/v1/twilight?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484")

func TestTwilightRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTwilightRouteObserver(t *testing.T) {
	// Build our expected observer section of body
	observer := gin.H{
		"datetime":  "2021-05-14T00:00:00Z",
		"latitude":  "19.798484",
		"longitude": "-155.468094",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the observer & whether or not it exists
	datetime, exists := response["observer"]["datetime"]
	assert.True(t, exists)

	// Grab the observer & whether or not it exists
	latitude, exists := response["observer"]["latitude"]
	assert.True(t, exists)

	// Grab the observer & whether or not it exists
	longitude, exists := response["observer"]["longitude"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, datetime, observer["datetime"])
	assert.Equal(t, latitude, observer["latitude"])
	assert.Equal(t, longitude, observer["longitude"])
}

func TestGetTwilightRouteAstronomicalTwilight(t *testing.T) {
	// Build our expected twilight section of body
	twilight := gin.H{
		"duration": 8.563931944444445,
		"from":     "2021-05-14T20:01:18-10:00",
		"location": "Pacific/Honolulu",
		"until":    "2021-05-15T04:35:08-10:00",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the twilight & whether or not it exists
	duration, exists := response["astronomical"]["duration"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	from, exists := response["astronomical"]["from"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	location, exists := response["astronomical"]["location"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	until, exists := response["astronomical"]["until"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, duration, twilight["duration"])
	assert.Equal(t, from, twilight["from"])
	assert.Equal(t, location, twilight["location"])
	assert.Equal(t, until, twilight["until"])
}

func TestGetTwilightRouteCivilTwilight(t *testing.T) {
	// Build our expected twilight section of body
	twilight := gin.H{
		"duration": 10.228770555555556,
		"from":     "2021-05-14T19:11:19-10:00",
		"location": "Pacific/Honolulu",
		"until":    "2021-05-15T05:25:03-10:00",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the twilight & whether or not it exists
	duration, exists := response["civil"]["duration"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	from, exists := response["civil"]["from"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	location, exists := response["civil"]["location"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	until, exists := response["civil"]["until"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, duration, twilight["duration"])
	assert.Equal(t, from, twilight["from"])
	assert.Equal(t, location, twilight["location"])
	assert.Equal(t, until, twilight["until"])
}

func TestGetTwilightRouteNauticalTwilight(t *testing.T) {
	// Build our expected twilight section of body
	twilight := gin.H{
		"duration": 9.402505555555557,
		"from":     "2021-05-14T19:36:08-10:00",
		"location": "Pacific/Honolulu",
		"until":    "2021-05-15T05:00:17-10:00",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the twilight & whether or not it exists
	duration, exists := response["nautical"]["duration"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	from, exists := response["nautical"]["from"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	location, exists := response["nautical"]["location"]
	assert.True(t, exists)

	// Grab the twilight & whether or not it exists
	until, exists := response["nautical"]["until"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, duration, twilight["duration"])
	assert.Equal(t, from, twilight["from"])
	assert.Equal(t, location, twilight["location"])
	assert.Equal(t, until, twilight["until"])
}
