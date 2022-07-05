package transit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Observer struct {
	Datetime  string  `json:"datetime"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Position struct {
	R   float64 `json:"R"`
	X   float64 `json:"X"`
	Alt float64 `json:"alt"`
	Az  float64 `json:"az"`
	Dec float64 `json:"dec"`
	Ra  float64 `json:"ra"`
}

type Properties struct {
	Maximum *string `json:"maximum"`
	Rise    *string `json:"rise"`
	Set     *string `json:"set"`
}

type Phase struct {
	Age          float64 `json:"age"`
	Angle        float64 `json:"angle"`
	D            float64 `json:"d"`
	Fraction     float64 `json:"fraction"`
	Illumination float64 `json:"illumination"`
	Separation   float64 `json:"separation"`
}

type Path struct {
	Datetime string  `json:"datetime"`
	Altitude float64 `json:"altitude"`
	Azimuth  float64 `json:"azimuth"`
	IsRise   bool    `json:"isRise"`
	IsSet    bool    `json:"isSet"`
}

type Response struct {
	Observer   Observer   `json:"observer"`
	Path       []Path     `json:"path"`
	Phase      Phase      `json:"phase"`
	Position   Position   `json:"position"`
	Properties Properties `json:"properties"`
}

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
var response Response

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Perform a GET request with that handler.
var w = performRequest(r, "GET", "/api/v1/transit?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=19.798484&ra=88.792958&dec=7.407064")

// Perform a GET request with that handler.
var x = performRequest(r, "GET", "/api/v1/transit?datetime=2021-05-14T00:00:00.000Z&longitude=-155.468094&latitude=45.798484&ra=88.792958&dec=-77.407064")

// Perform a GET request with that handler.
var y = performRequest(r, "GET", "/api/v1/transit?datetime=2021-05-14T06:52:13.000Z&longitude=-155.468094&latitude=19.798484&ra=88.792958&dec=7.407064")

func TestTransitRouteStatusCode(t *testing.T) {
	// Assert we encoded correctly, the request gives a 200:
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransitRouteMoon(t *testing.T) {
	// Build our expected observer section of body
	phase := gin.H{
		"age":          1.2222287803073832,
		"angle":        156.46390817398918,
		"d":            23.47659745538946,
		"fraction":     0.041388566239529356,
		"illumination": 4.1595644017041575,
		"separation":   20.18056657827112,
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the phase & whether or not it exists
	age := response.Phase.Age

	// Grab the phase & whether or not it exists
	angle := response.Phase.Angle

	// Grab the phase & whether or not it exists
	d := response.Phase.D

	// Grab the phase & whether or not it exists
	fraction := response.Phase.Fraction

	// Grab the phase & whether or not it exists
	illumination := response.Phase.Illumination

	// Grab the phase & whether or not it exists
	separation := response.Phase.Separation

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
		"latitude":  19.798484,
		"longitude": -155.468094,
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the observer & whether or not it exists
	datetime := response.Observer.Datetime

	// Grab the observer & whether or not it exists
	latitude := response.Observer.Latitude

	// Grab the observer & whether or not it exists
	longitude := response.Observer.Longitude

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, datetime, observer["datetime"])
	assert.Equal(t, latitude, observer["latitude"])
	assert.Equal(t, longitude, observer["longitude"])
}

func TestGetTransitRoutePosition(t *testing.T) {
	// Build our expected position section of body
	position := gin.H{
		"R":   0.005219234547163293,
		"X":   1.0465576817848306,
		"alt": 72.80058854788766,
		"az":  134.39667229414232,
		"dec": 7.407064,
		"ra":  88.792958,
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the position & whether or not it exists
	alt := response.Position.Alt

	// Grab the position & whether or not it exists
	az := response.Position.Az

	// Grab the position & whether or not it exists
	ra := response.Position.Ra

	// Grab the position & whether or not it exists
	dec := response.Position.Dec

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, alt, position["alt"])
	assert.Equal(t, az, position["az"])
	assert.Equal(t, ra, position["ra"])
	assert.Equal(t, dec, position["dec"])
}

func TestGetTransitRouteProperties(t *testing.T) {
	// Build our expected properties section of transit
	properties := gin.H{
		"maximum": "2021-05-14T12:39:25-10:00",
		"rise":    "2021-05-14T08:35:25-10:00",
		"set":     "2021-05-14T20:54:51-10:00",
	}

	// Convert the JSON response:
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the properties & whether or not it exists
	rise := *response.Properties.Rise

	// Grab the properties & whether or not it exists
	set := *response.Properties.Set

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, rise, properties["rise"])
	assert.Equal(t, set, properties["set"])
}

func TestGetTransitRoutePropertiesAlt(t *testing.T) {
	// Build our expected transit section of body
	properties := gin.H{
		"maximum": "2021-05-14T12:39:25-10:00",
		"rise":    "2021-05-14T08:35:25-10:00",
		"set":     "2021-05-14T20:54:51-10:00",
	}

	// Convert the JSON response:
	err := json.Unmarshal(y.Body.Bytes(), &response)

	// Grab the properties & whether or not it exists
	rise := *response.Properties.Rise

	fmt.Println(rise)

	// Grab the properties & whether or not it exists
	set := *response.Properties.Set

	fmt.Println(set)

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Equal(t, rise, properties["rise"])
	assert.Equal(t, set, properties["set"])
}

func TestGetTransitRoutePropertiesNotAboveHorizon(t *testing.T) {
	// Convert the JSON response:
	err := json.Unmarshal(x.Body.Bytes(), &response)

	// Grab the properties & whether or not it exists
	maximum := response.Properties.Maximum

	// Grab the properties & whether or not it exists
	rise := response.Properties.Rise

	// Grab the properties & whether or not it exists
	set := response.Properties.Set

	// Assert on the correctness of the response:
	assert.Nil(t, err)
	assert.Nil(t, maximum)
	assert.Nil(t, rise)
	assert.Nil(t, set)
}
