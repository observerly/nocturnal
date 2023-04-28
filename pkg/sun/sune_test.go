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

func SetupSuneRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	r.GET("/api/v2/sun", GetSun)

	return r
}

// Setup the Gin API router:
var sr = SetupSuneRouter()

func performSuneRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var sw = performSuneRequest(sr, "GET", "/api/v2/sun?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484")

func TestSuneRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, sw.Code)
}

func TestGetSuneRouteObserver(t *testing.T) {
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

func TestGetSunRouteRise(t *testing.T) {
	// Build our expected rise section of body
	rise := gin.H{
		"LCT": "2021-05-14T05:49:45-10:00",
		"UTC": "2021-05-14T15:49:45Z",
		"alt": 1.3101013887429336,
		"az":  70.47433002623417,
		"dec": 18.792075895178936,
		"ra":  51.71630525455092,
	}

	// Convert the JSON response:
	err := json.Unmarshal(sw.Body.Bytes(), &response)

	// Obtain the Local Civil Time rise of the rise and test whether or not it exists:
	LCT, exists := response["rise"]["LCT"]
	assert.True(t, exists)

	// Obtain the Universal Time rise of the rise and test whether or not it exists:
	UTC, exists := response["rise"]["UTC"]
	assert.True(t, exists)

	// Obtain the altitude of the rise and test whether or not it exists:
	alt, exists := response["rise"]["alt"]
	assert.True(t, exists)

	// Obtain the azimuth of the rise and test whether or not it exists:
	az, exists := response["rise"]["az"]
	assert.True(t, exists)

	// Obtain the declination of the rise and test whether or not it exists:
	dec, exists := response["rise"]["dec"]
	assert.True(t, exists)

	// Obtain the right ascension of the rise and test whether or not it exists:
	ra, exists := response["rise"]["ra"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, LCT, rise["LCT"])
	assert.Equal(t, UTC, rise["UTC"])
	assert.InDelta(t, alt, rise["alt"], precision)
	assert.InDelta(t, az, rise["az"], precision)
	assert.InDelta(t, dec, rise["dec"], precision)
	assert.InDelta(t, ra, rise["ra"], precision)
}

func TestGetSunRouteSet(t *testing.T) {
	// Build our expected set section of body
	set := gin.H{
		"LCT": "2021-05-14T18:46:50-10:00",
		"UTC": "2021-05-15T04:46:50Z",
		"alt": -3.2429444558408953,
		"az":  68.55790252346988,
		"dec": 18.919575832754642,
		"ra":  52.24955492872896,
	}

	// Convert the JSON response:
	err := json.Unmarshal(sw.Body.Bytes(), &response)

	// Obtain the Local Civil Time set of the set and test whether or not it exists:
	LCT, exists := response["set"]["LCT"]
	assert.True(t, exists)

	// Obtain the Universal Time set of the set and test whether or not it exists:
	UTC, exists := response["set"]["UTC"]
	assert.True(t, exists)

	// Obtain the altitude of the set and test whether or not it exists:
	alt, exists := response["set"]["alt"]
	assert.True(t, exists)

	// Obtain the azimuth of the set and test whether or not it exists:
	az, exists := response["set"]["az"]
	assert.True(t, exists)

	// Obtain the declination of the set and test whether or not it exists:
	dec, exists := response["set"]["dec"]
	assert.True(t, exists)

	// Obtain the right ascension of the set and test whether or not it exists:
	ra, exists := response["set"]["ra"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, LCT, set["LCT"])
	assert.Equal(t, UTC, set["UTC"])
	assert.InDelta(t, alt, set["alt"], precision)
	assert.InDelta(t, az, set["az"], precision)
	assert.InDelta(t, dec, set["dec"], precision)
	assert.InDelta(t, ra, set["ra"], precision)
}
