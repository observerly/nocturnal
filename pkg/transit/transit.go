package transit

import (
	"fmt"
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

	ra := c.DefaultQuery("ra", strconv.Itoa(8))

	dec := c.DefaultQuery("dec", strconv.Itoa(8))

	lon := c.DefaultQuery("longitude", strconv.Itoa(8))

	lat := c.DefaultQuery("latitude", strconv.Itoa(8))

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

	phase := gin.H{
		"age":          fmt.Sprintf("%f", mph.Days),
		"angle":        fmt.Sprintf("%f", mph.Angle),
		"d":            fmt.Sprintf("%f", mph.Age),
		"fraction":     fmt.Sprintf("%f", mph.Fraction),
		"illumination": fmt.Sprintf("%f", mph.Illumination),
		"separation":   fmt.Sprintf("%f", separation),
	}

	observer := gin.H{
		"datetime":  datetime,
		"longitude": fmt.Sprintf("%f", longitude),
		"latitude":  fmt.Sprintf("%f", latitude),
	}

	position := gin.H{
		"alt": fmt.Sprintf("%f", hz.Altitude),
		"az":  fmt.Sprintf("%f", hz.Azimuth),
		"ra":  fmt.Sprintf("%f", rightAscension),
		"dec": fmt.Sprintf("%f", declination),
		"R":   fmt.Sprintf("%v", refraction),
		"X":   fmt.Sprintf("%v", airmass),
	}

	transit := gin.H{
		"maximum": utils.FormatDatetimeRFC3339(tr.Maximum),
		"rise":    utils.FormatDatetimeRFC3339(tr.Rise),
		"set":     utils.FormatDatetimeRFC3339(tr.Set),
		"path":    path,
	}

	c.JSON(http.StatusOK, gin.H{
		"phase":    phase,
		"observer": observer,
		"position": position,
		"transit":  transit,
	})
}
