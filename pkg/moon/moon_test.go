package moon

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupMoonRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	r.GET("/api/v1/moon", GetMoon)

	return r
}

// Setup the Gin API router:
var r = SetupMoonRouter()

// Setup the base response struct:
var response map[string]map[string]string

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var w = performRequest(r, "GET", "/api/v1/moon?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484")

func TestMoonRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetMoonRouteObserver(t *testing.T) {
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

func TestGetMoonRoutePhase(t *testing.T) {
	// Build our expected phase section of body
	phase := gin.H{
		"age":          "1.222229",
		"angle":        "156.463908",
		"d":            "23.476597",
		"fraction":     "0.041389",
		"illumination": "4.159564",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the phase & whether or not it exists
	age, exists := response["phase"]["age"]
	assert.True(t, exists)

	// Grab the phase & whether or not it exists
	angle, exists := response["phase"]["angle"]
	assert.True(t, exists)

	// Grab the phase & whether or not it exists
	d, exists := response["phase"]["d"]
	assert.True(t, exists)

	// Grab the phase & whether or not it exists
	fraction, exists := response["phase"]["fraction"]
	assert.True(t, exists)

	// Grab the phase & whether or not it exists
	illumination, exists := response["phase"]["illumination"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, age, phase["age"])
	assert.Equal(t, angle, phase["angle"])
	assert.Equal(t, d, phase["d"])
	assert.Equal(t, fraction, phase["fraction"])
	assert.Equal(t, illumination, phase["illumination"])
}

func TestGetMoonRoutePosition(t *testing.T) {
	// Build our expected position section of body
	position := gin.H{
		"alt": "86.192481",
		"az":  "3.481574",
		"dec": "23.598793",
		"ra":  "76.239624",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the position & whether or not it exists
	alt, exists := response["position"]["alt"]
	assert.True(t, exists)

	// Grab the position & whether or not it exists
	az, exists := response["position"]["az"]
	assert.True(t, exists)

	// Grab the position & whether or not it exists
	ra, exists := response["position"]["ra"]
	assert.True(t, exists)

	// Grab the position & whether or not it exists
	dec, exists := response["position"]["dec"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, alt, position["alt"])
	assert.Equal(t, az, position["az"])
	assert.Equal(t, ra, position["ra"])
	assert.Equal(t, dec, position["dec"])
}
