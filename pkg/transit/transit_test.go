package transit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTransitRouter() *gin.Engine {
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	r.GET("/api/v2/transit", GetTransit)

	return r
}

// Setup the Gin API router:
var r = SetupTransitRouter()

// Setup the base response struct:
var response struct {
	Observer map[string]interface{} `json:"observer"`
	Rise     map[string]interface{} `json:"rise"`
	Set      map[string]interface{} `json:"set"`
	Maximum  map[string]interface{} `json:"maximum"`
	Path     []interface{}          `json:"path"`
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var w = performRequest(r, "GET", "/api/v2/transit?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484&ra=88.792958&dec=7.407064")

// Perform a GET request with that handler.
var x = performRequest(r, "GET", "/api/v2/transit?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=45.798484&ra=88.792958&dec=-77.407064")

// Perform a GET request with that handler.
var y = performRequest(r, "GET", "/api/v2/transit?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484&ra=88.792958&dec=77.407064")

func TestTransitRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransitRouteObserver(t *testing.T) {
	// Build our expected observer section of body
	observer := gin.H{
		"datetime":  "2021-05-14T00:00:00Z",
		"latitude":  19.798484,
		"longitude": -155.468094,
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the observer & whether or not it exists
	datetime, exists := response.Observer["datetime"]
	assert.True(t, exists)

	// Grab the observer & whether or not it exists
	latitude, exists := response.Observer["latitude"]
	assert.True(t, exists)

	// Grab the observer & whether or not it exists
	longitude, exists := response.Observer["longitude"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, datetime, observer["datetime"])
	assert.Equal(t, latitude, observer["latitude"])
	assert.Equal(t, longitude, observer["longitude"])
}

func TestGetTransitRouteRise(t *testing.T) {
	// Build our expected rise section of body
	rise := gin.H{
		"LCT":          "2021-05-14T08:35:25-10:00",
		"R":            0.31247646372444293,
		"UTC":          "2021-05-14T18:35:25Z",
		"X":            22.21853513271382,
		"age":          2.185508160545753,
		"alt":          1.5727806816314758,
		"angle":        148.34943756923096,
		"az":           82.69308817455995,
		"dec":          7.407064,
		"fraction":     0.07400827956860168,
		"illumination": 7.436790249940451,
		"ra":           88.792958,
		"separation":   17.489602973521922,
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Obtain the Local Civil Time rise of the rise and test whether or not it exists:
	LCT, exists := response.Rise["LCT"]
	assert.True(t, exists)

	// Obtain the refraction of the rise and test whether or not it exists:
	R, exists := response.Rise["R"]
	assert.True(t, exists)

	// Obtain the Universal Time rise of the rise and test whether or not it exists:
	UTC, exists := response.Rise["UTC"]
	assert.True(t, exists)

	// Obtain the airmass (X) of the rise and test whether or not it exists:
	X, exists := response.Rise["X"]
	assert.True(t, exists)

	// Obtain the age at the the rise and test whether or not it exists:
	age, exists := response.Rise["age"]
	assert.True(t, exists)

	// Obtain the altitude of the rise and test whether or not it exists:
	alt, exists := response.Rise["alt"]
	assert.True(t, exists)

	// Obtain the angle of the rise and test whether or not it exists:
	angle, exists := response.Rise["angle"]
	assert.True(t, exists)

	// Obtain the azimuth of the rise and test whether or not it exists:
	az, exists := response.Rise["az"]
	assert.True(t, exists)

	// Obtain the declination of the rise and test whether or not it exists:
	dec, exists := response.Rise["dec"]
	assert.True(t, exists)

	// Obtain the fraction of the rise and test whether or not it exists:
	fraction, exists := response.Rise["fraction"]
	assert.True(t, exists)

	// Obtain the illumination of the rise and test whether or not it exists:
	illumination, exists := response.Rise["illumination"]
	assert.True(t, exists)

	// Obtain the right ascension of the rise and test whether or not it exists:
	ra, exists := response.Rise["ra"]
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

func TestGetTransitRouteMaximum(t *testing.T) {
	// Build our expected maximum section of body
	maximum := gin.H{
		"LCT":          "2021-05-14T12:39:25-10:00",
		"R":            0.010331651302290006,
		"UTC":          "2021-05-14T22:39:25Z",
		"X":            1.1715008193883227,
		"age":          2.467644467039966,
		"alt":          58.54930192457176,
		"angle":        146.31204868601841,
		"az":           109.02539763910731,
		"dec":          7.407064,
		"fraction":     0.08356231511256518,
		"illumination": 8.396460922585103,
		"ra":           88.792958,
		"separation":   17.52318422105279,
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Obtain the Local Civil Time maximum of the maximum and test whether or not it exists:
	LCT, exists := response.Maximum["LCT"]
	assert.True(t, exists)

	// Obtain the refraction of the maximum and test whether or not it exists:
	R, exists := response.Maximum["R"]
	assert.True(t, exists)

	// Obtain the Universal Time maximum of the maximum and test whether or not it exists:
	UTC, exists := response.Maximum["UTC"]
	assert.True(t, exists)

	// Obtain the airmass (X) of the maximum and test whether or not it exists:
	X, exists := response.Maximum["X"]
	assert.True(t, exists)

	// Obtain the age at the the maximum and test whether or not it exists:
	age, exists := response.Maximum["age"]
	assert.True(t, exists)

	// Obtain the altitude of the maximum and test whether or not it exists:
	alt, exists := response.Maximum["alt"]
	assert.True(t, exists)

	// Obtain the angle of the maximum and test whether or not it exists:
	angle, exists := response.Maximum["angle"]
	assert.True(t, exists)

	// Obtain the azimuth of the maximum and test whether or not it exists:
	az, exists := response.Maximum["az"]
	assert.True(t, exists)

	// Obtain the declination of the maximum and test whether or not it exists:
	dec, exists := response.Maximum["dec"]
	assert.True(t, exists)

	// Obtain the fraction of the maximum and test whether or not it exists:
	fraction, exists := response.Maximum["fraction"]
	assert.True(t, exists)

	// Obtain the illumination of the maximum and test whether or not it exists:
	illumination, exists := response.Maximum["illumination"]
	assert.True(t, exists)

	// Obtain the right ascension of the maximum and test whether or not it exists:
	ra, exists := response.Maximum["ra"]
	assert.True(t, exists)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, R, maximum["R"])
	assert.Equal(t, LCT, maximum["LCT"])
	assert.Equal(t, UTC, maximum["UTC"])
	assert.Equal(t, X, maximum["X"])
	assert.Equal(t, age, maximum["age"])
	assert.Equal(t, alt, maximum["alt"])
	assert.Equal(t, angle, maximum["angle"])
	assert.Equal(t, az, maximum["az"])
	assert.Equal(t, dec, maximum["dec"])
	assert.Equal(t, fraction, maximum["fraction"])
	assert.Equal(t, illumination, maximum["illumination"])
	assert.Equal(t, ra, maximum["ra"])
}

func TestGetTransitRouteSet(t *testing.T) {
	// Build our expected set section of body
	set := gin.H{
		"LCT":          "2021-05-14T20:54:51-10:00",
		"R":            nil,
		"UTC":          "2021-05-15T06:54:51Z",
		"X":            nil,
		"age":          3.0891483187381277,
		"alt":          -1.8540744329511798,
		"angle":        142.16654610640515,
		"az":           81.44596647698675,
		"dec":          7.407064,
		"fraction":     0.10460841854965064,
		"illumination": 10.51014934125243,
		"ra":           88.792958,
		"separation":   18.281831615702153,
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Obtain the Local Civil Time set of the set and test whether or not it exists:
	LCT, exists := response.Set["LCT"]
	assert.True(t, exists)

	// Obtain the refraction of the set and test whether or not it exists:
	R, exists := response.Set["R"]
	assert.True(t, exists)

	// Obtain the Universal Time set of the set and test whether or not it exists:
	UTC, exists := response.Set["UTC"]
	assert.True(t, exists)

	// Obtain the airmass (X) of the set and test whether or not it exists:
	X, exists := response.Set["X"]
	assert.True(t, exists)

	// Obtain the age at the the set and test whether or not it exists:
	age, exists := response.Set["age"]
	assert.True(t, exists)

	// Obtain the altitude of the set and test whether or not it exists:
	alt, exists := response.Set["alt"]
	assert.True(t, exists)

	// Obtain the angle of the set and test whether or not it exists:
	angle, exists := response.Set["angle"]
	assert.True(t, exists)

	// Obtain the azimuth of the set and test whether or not it exists:
	az, exists := response.Set["az"]
	assert.True(t, exists)

	// Obtain the declination of the set and test whether or not it exists:
	dec, exists := response.Set["dec"]
	assert.True(t, exists)

	// Obtain the fraction of the set and test whether or not it exists:
	fraction, exists := response.Set["fraction"]
	assert.True(t, exists)

	// Obtain the illumination of the set and test whether or not it exists:
	illumination, exists := response.Set["illumination"]
	assert.True(t, exists)

	// Obtain the right ascension of the set and test whether or not it exists:
	ra, exists := response.Set["ra"]
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

func TestGetTransitRouteAlwaysBelowHorizon(t *testing.T) {
	// Convert the JSON response:
	err := json.Unmarshal(x.Body.Bytes(), &response)

	// Assert on the correctness of the response:
	assert.Nil(t, err)

	// Assert that the response is empty:
	assert.Equal(t, 0, len(response.Set))
	assert.Equal(t, 0, len(response.Maximum))
	assert.Equal(t, 0, len(response.Rise))
}

func TestGetTransitRouteAlwaysAboveHorizon(t *testing.T) {
	maximum := gin.H{}

	// Convert the JSON response:
	err := json.Unmarshal(y.Body.Bytes(), &response)

	// Assert on the correctness of the response:
	assert.Nil(t, err)

	// Assert that the response is empty:
	assert.Equal(t, 0, len(response.Set))
	assert.Equal(t, maximum, response.Maximum)
	assert.Equal(t, 0, len(response.Rise))
}
