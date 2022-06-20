package transit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/observerly/dusk/pkg/dusk"
	"github.com/observerly/nocturnal/internal/utils"
	"github.com/stretchr/testify/assert"
)

func SetupTransitRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	r.GET("/api/v1/transit", GetTransit)

	return r
}

// Setup the Gin API router:
var r = SetupTransitRouter()

// Setup the base response struct:
var response map[string]map[string]interface{}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var w = performRequest(r, "GET", "/api/v1/transit?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484&ra=88.792958&dec=7.407064")

// Perform a GET request with that handler.
var x = performRequest(r, "GET", "/api/v1/transit?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484&ra=88.792958&dec=-77.407064")

func TestTransitRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransitRouteMoon(t *testing.T) {
	// Build our expected observer section of body
	phase := gin.H{
		"age":          "1.222229",
		"angle":        "156.463908",
		"d":            "23.476597",
		"fraction":     "0.041389",
		"illumination": "4.159564",
		"separation":   "20.180567",
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

	// Grab the phase & whether or not it exists
	separation, exists := response["phase"]["separation"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, age, phase["age"])
	assert.Equal(t, angle, phase["angle"])
	assert.Equal(t, d, phase["d"])
	assert.Equal(t, fraction, phase["fraction"])
	assert.Equal(t, illumination, phase["illumination"])
	assert.Equal(t, separation, phase["separation"])
}

func TestGetTransitRouteObserver(t *testing.T) {
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

func TestGetTransitRoutePosition(t *testing.T) {
	// Build our expected position section of body
	position := gin.H{
		"R":   "0.005219",
		"X":   "1.046558",
		"alt": "72.800589",
		"az":  "134.396672",
		"dec": "7.407064",
		"ra":  "88.792958",
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

func TestGetTransitRouteTransit(t *testing.T) {
	datetime, _ := utils.ParseDatetimeRFC3339("2021-05-14T00:00:00.000Z")

	eq := dusk.EquatorialCoordinate{RightAscension: 88.792958, Declination: 7.407064}

	coordinates, _ := dusk.GetObjectHorizontalCoordinatesForDay(datetime, eq, -155.468094, 19.798484)

	// Build our expected transit section of body
	transit := gin.H{
		"maximum": "2021-05-14T12:39:25-10:00",
		"rise":    "2021-05-14T08:35:25-10:00",
		"set":     "2021-05-14T20:54:51-10:00",
		"path":    coordinates,
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

func TestGetTransitRouteTransitNotAboveHorizon(t *testing.T) {
	// Build our expected transit section of body
	transit := gin.H{
		"maximum": nil,
		"rise":    nil,
		"set":     nil,
	}

	// Convert the JSON response:
	err := json.Unmarshal(x.Body.Bytes(), &response)

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
