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

func SetupLuneRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	r.GET("/api/v2/moon", GetMoon)

	return r
}

// Setup the Gin API router:
var lr = SetupLuneRouter()

func performLuneRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var lw = performLuneRequest(lr, "GET", "/api/v2/moon?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484")

func TestLuneRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, lw.Code)
}

func TestGetLuneRouteObserver(t *testing.T) {
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

func TestGetMoonRouteRise(t *testing.T) {
	// Build our expected rise section of body
	rise := gin.H{
		"LCT":          "2021-05-14T07:57:00-10:00",
		"R":            0.45726609910544613,
		"UTC":          "2021-05-14T17:57:00Z",
		"X":            35.82232848697485,
		"age":          2.142546923777335,
		"alt":          0.18326186176169362,
		"angle":        148.6700291081809,
		"az":           63.73505782242811,
		"dec":          24.673459614318453,
		"fraction":     0.07255347501615561,
		"illumination": 7.290652119635177,
		"ra":           85.78031360389217,
	}

	// Convert the JSON response:
	err := json.Unmarshal(lw.Body.Bytes(), &response)

	// Obtain the Local Civil Time rise of the rise and test whether or not it exists:
	LCT, exists := response["rise"]["LCT"]
	assert.True(t, exists)

	// Obtain the refraction of the rise and test whether or not it exists:
	R, exists := response["rise"]["R"]
	assert.True(t, exists)

	// Obtain the Universal Time rise of the rise and test whether or not it exists:
	UTC, exists := response["rise"]["UTC"]
	assert.True(t, exists)

	// Obtain the airmass (X) of the rise and test whether or not it exists:
	X, exists := response["rise"]["X"]
	assert.True(t, exists)

	// Obtain the age at the the rise and test whether or not it exists:
	age, exists := response["rise"]["age"]
	assert.True(t, exists)

	// Obtain the altitude of the rise and test whether or not it exists:
	alt, exists := response["rise"]["alt"]
	assert.True(t, exists)

	// Obtain the angle of the rise and test whether or not it exists:
	angle, exists := response["rise"]["angle"]
	assert.True(t, exists)

	// Obtain the azimuth of the rise and test whether or not it exists:
	az, exists := response["rise"]["az"]
	assert.True(t, exists)

	// Obtain the declination of the rise and test whether or not it exists:
	dec, exists := response["rise"]["dec"]
	assert.True(t, exists)

	// Obtain the fraction of the rise and test whether or not it exists:
	fraction, exists := response["rise"]["fraction"]
	assert.True(t, exists)

	// Obtain the illumination of the rise and test whether or not it exists:
	illumination, exists := response["rise"]["illumination"]
	assert.True(t, exists)

	// Obtain the right ascension of the rise and test whether or not it exists:
	ra, exists := response["rise"]["ra"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, R, rise["R"])
	assert.Equal(t, LCT, rise["LCT"])
	assert.Equal(t, UTC, rise["UTC"])
	assert.Equal(t, X, rise["X"])
	assert.Equal(t, age, rise["age"])
	assert.Equal(t, alt, rise["alt"])
	assert.Equal(t, angle, rise["angle"])
	assert.Equal(t, az, rise["az"])
	assert.Equal(t, dec, rise["dec"])
	assert.Equal(t, fraction, rise["fraction"])
	assert.Equal(t, illumination, rise["illumination"])
	assert.Equal(t, ra, rise["ra"])
}

func TestGetMoonRouteSet(t *testing.T) {
	// Build our expected set section of body
	set := gin.H{
		"LCT":          "2021-05-14T21:42:00-10:00",
		"R":            nil,
		"UTC":          "2021-05-15T07:42:00Z",
		"X":            nil,
		"age":          3.151614688448534,
		"alt":          -0.1253140634039632,
		"angle":        141.7714490563052,
		"az":           62.95233340505419,
		"dec":          25.284322055858368,
		"fraction":     0.10672372913810846,
		"illumination": 10.722568110680397,
		"ra":           93.35313596587092,
	}

	// Convert the JSON response:
	err := json.Unmarshal(lw.Body.Bytes(), &response)

	// Obtain the Local Civil Time set of the set and test whether or not it exists:
	LCT, exists := response["set"]["LCT"]
	assert.True(t, exists)

	// Obtain the refraction of the set and test whether or not it exists:
	R, exists := response["set"]["R"]
	assert.True(t, exists)

	// Obtain the Universal Time set of the set and test whether or not it exists:
	UTC, exists := response["set"]["UTC"]
	assert.True(t, exists)

	// Obtain the airmass (X) of the set and test whether or not it exists:
	X, exists := response["set"]["X"]
	assert.True(t, exists)

	// Obtain the age at the the set and test whether or not it exists:
	age, exists := response["set"]["age"]
	assert.True(t, exists)

	// Obtain the altitude of the set and test whether or not it exists:
	alt, exists := response["set"]["alt"]
	assert.True(t, exists)

	// Obtain the angle of the set and test whether or not it exists:
	angle, exists := response["set"]["angle"]
	assert.True(t, exists)

	// Obtain the azimuth of the set and test whether or not it exists:
	az, exists := response["set"]["az"]
	assert.True(t, exists)

	// Obtain the declination of the set and test whether or not it exists:
	dec, exists := response["set"]["dec"]
	assert.True(t, exists)

	// Obtain the fraction of the set and test whether or not it exists:
	fraction, exists := response["set"]["fraction"]
	assert.True(t, exists)

	// Obtain the illumination of the set and test whether or not it exists:
	illumination, exists := response["set"]["illumination"]
	assert.True(t, exists)

	// Obtain the right ascension of the set and test whether or not it exists:
	ra, exists := response["set"]["ra"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, R, set["R"])
	assert.Equal(t, LCT, set["LCT"])
	assert.Equal(t, UTC, set["UTC"])
	assert.Equal(t, X, set["X"])
	assert.Equal(t, age, set["age"])
	assert.Equal(t, alt, set["alt"])
	assert.Equal(t, angle, set["angle"])
	assert.Equal(t, az, set["az"])
	assert.Equal(t, dec, set["dec"])
	assert.Equal(t, fraction, set["fraction"])
	assert.Equal(t, illumination, set["illumination"])
	assert.Equal(t, ra, set["ra"])
}
