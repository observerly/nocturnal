package transit

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/observerly/dusk/pkg/dusk"
	"github.com/observerly/nocturnal/internal/utils"
)

// GET Transit
func GetTransit(c *gin.Context) {
	d := c.DefaultQuery("datetime", time.Now().String())

	ra := c.DefaultQuery("ra", strconv.Itoa(0))

	dec := c.DefaultQuery("dec", strconv.Itoa(0))

	lon := c.DefaultQuery("longitude", strconv.Itoa(0))

	lat := c.DefaultQuery("latitude", strconv.Itoa(0))

	datetime, _ := utils.ParseDatetimeRFC3339(d)

	longitude, _ := strconv.ParseFloat(lon, 64)

	latitude, _ := strconv.ParseFloat(lat, 64)

	rightAscension, _ := strconv.ParseFloat(ra, 64)

	declination, _ := strconv.ParseFloat(dec, 64)

	eq := dusk.EquatorialCoordinate{RightAscension: rightAscension, Declination: declination}

	hz := dusk.ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, eq)

	mec := dusk.GetLunarEclipticPosition(datetime)

	meq := dusk.GetLunarEquatorialPosition(datetime)

	mph := dusk.GetLunarPhase(datetime, longitude, mec)

	tr, _ := dusk.GetObjectTransit(datetime, eq, latitude, longitude)

	path, _ := dusk.GetObjectHorizontalCoordinatesForDay(datetime, eq, longitude, latitude)

	airmass := dusk.GetRelativeAirMass(hz.Altitude)

	refraction := dusk.GetAtmosphericRefraction(hz.Altitude)

	separation := dusk.GetAngularSeparation(dusk.Coordinate{Latitude: eq.Declination, Longitude: eq.RightAscension}, dusk.Coordinate{Latitude: meq.Declination, Longitude: meq.RightAscension})

	observer := gin.H{
		"datetime":  datetime,
		"longitude": longitude,
		"latitude":  latitude,
	}

	phase := gin.H{
		"age":          mph.Days,
		"angle":        mph.Angle,
		"d":            mph.Age,
		"fraction":     mph.Fraction,
		"illumination": mph.Illumination,
		"separation":   separation,
	}

	position := gin.H{
		"alt": hz.Altitude,
		"az":  hz.Azimuth,
		"ra":  rightAscension,
		"dec": declination,
		"R":   refraction,
		"X":   airmass,
	}

	properties := gin.H{
		"maximum": utils.FormatDatetimeRFC3339(tr.Maximum),
		"rise":    utils.FormatDatetimeRFC3339(tr.Rise),
		"set":     utils.FormatDatetimeRFC3339(tr.Set),
	}

	c.JSON(http.StatusOK, gin.H{
		"phase":      phase,
		"observer":   observer,
		"position":   position,
		"properties": properties,
		"path":       path,
	})
}
