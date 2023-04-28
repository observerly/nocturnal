package sun

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupSunRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	r.GET("/api/v1/sun", GetSunDeprecatedV1)

	return r
}

// Setup the Gin API router:
var r = SetupSunRouter()

// Setup the base response struct:
var response map[string]map[string]interface{}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var w = performRequest(r, "GET", "/api/v1/sun?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484")

func TestSunRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)
}

var precision = 0.0000001

func TestGetSunRouteObserver(t *testing.T) {
	// Build our expected observer section of body
	observer := gin.H{
		"datetime":  "2021-05-14T00:00:00Z",
		"latitude":  19.798484,
		"longitude": -155.468094,
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

func TestGetSunRoutePosition(t *testing.T) {
	// Build our expected position section of body
	position := gin.H{
		"alt": 65.98487307697896,
		"az":  88.4839666699854,
		"dec": 18.634152331055457,
		"ra":  51.065497132296336,
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
	assert.InDelta(t, alt, position["alt"], precision)
	assert.InDelta(t, az, position["az"], precision)
	assert.InDelta(t, ra, position["ra"], precision)
	assert.InDelta(t, dec, position["dec"], precision)
}

func TestGetSunRouteTransit(t *testing.T) {
	// Build our expected transit section of body
	transit := gin.H{
		"rise": "2021-05-14T05:49:45-10:00",
		"set":  "2021-05-14T18:46:50-10:00",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the transit & whether or not it exists
	rise, exists := response["transit"]["rise"]
	assert.True(t, exists)

	// Grab the transit & whether or not it exists
	set, exists := response["transit"]["set"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)

	assert.Equal(t, rise, transit["rise"])
	assert.Equal(t, set, transit["set"])
}

func TestGetSunRouteTomorrow(t *testing.T) {
	// Build our expected transit section of body
	tomorrow := gin.H{
		"rise": "2021-05-15T05:49:30-10:00",
		"set":  "2021-05-15T18:47:08-10:00",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the tomorrow & whether or not it exists
	rise, exists := response["tomorrow"]["rise"]
	assert.True(t, exists)

	// Grab the tomorrow & whether or not it exists
	set, exists := response["tomorrow"]["set"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, rise, tomorrow["rise"])
	assert.Equal(t, set, tomorrow["set"])
}
