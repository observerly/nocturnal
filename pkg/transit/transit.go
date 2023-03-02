package transit

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/observerly/dusk/pkg/dusk"
	"github.com/observerly/nocturnal/internal/query"
	"github.com/observerly/nocturnal/internal/utils"
)

// GET /transit
func GetTransitDeprecatedV1(c *gin.Context) {
	d, lon, lat := query.GetDefaultObserverParams(c)

	ra := c.DefaultQuery("ra", strconv.Itoa(0))

	dec := c.DefaultQuery("dec", strconv.Itoa(0))

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

func GetStandardTransitProperties(datetime *time.Time, eq dusk.EquatorialCoordinate, longitude float64, latitude float64) gin.H {
	if datetime == nil {
		return nil
	}

	ec := dusk.GetLunarEclipticPositionLawrence(datetime.UTC())

	hz := dusk.ConvertEquatorialCoordinateToHorizontal(datetime.UTC(), longitude, latitude, eq)

	ph := dusk.GetLunarPhase(datetime.UTC(), longitude, ec)

	airmass := dusk.GetRelativeAirMass(hz.Altitude)

	refraction := dusk.GetAtmosphericRefraction(hz.Altitude)

	mec := dusk.GetLunarEclipticPositionLawrence(datetime.UTC())

	meq := dusk.ConvertEclipticCoordinateToEquatorial(datetime.UTC(), mec)

	separation := dusk.GetAngularSeparation(dusk.Coordinate{Latitude: eq.Declination, Longitude: eq.RightAscension}, dusk.Coordinate{Latitude: meq.Declination, Longitude: meq.RightAscension})

	return gin.H{
		"UTC":          datetime.UTC().Format(time.RFC3339),
		"LCT":          datetime.Format(time.RFC3339),
		"alt":          hz.Altitude,
		"az":           hz.Azimuth,
		"ra":           eq.RightAscension,
		"dec":          eq.Declination,
		"age":          ph.Days,
		"angle":        ph.Angle,
		"fraction":     ph.Fraction,
		"illumination": ph.Illumination,
		"R":            refraction,
		"X":            airmass,
		"separation":   separation,
	}
}

// GET /transit v2
func GetTransit(c *gin.Context) {
	d, lon, lat := query.GetDefaultObserverParams(c)

	datetime, _ := utils.ParseDatetimeRFC3339(d)

	longitude, _ := strconv.ParseFloat(lon, 64)

	latitude, _ := strconv.ParseFloat(lat, 64)

	// Create the Observer gin.H JSON object representation:
	observer := gin.H{
		"datetime":  datetime,
		"longitude": longitude,
		"latitude":  latitude,
	}

	// Parse the Right Ascension from the request query:
	ra, err := strconv.ParseFloat(c.DefaultQuery("ra", strconv.Itoa(0)), 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parse the Declination from the request query:
	dec, err := strconv.ParseFloat(c.DefaultQuery("dec", strconv.Itoa(0)), 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	eq := dusk.EquatorialCoordinate{
		RightAscension: ra,
		Declination:    dec,
	}

	// Get the transit times:
	transit, _ := dusk.GetObjectTransit(datetime, eq, latitude, longitude)

	// Create the Rise gin.H JSON object representation:
	rise := GetStandardTransitProperties(transit.Rise, eq, longitude, latitude)

	if transit.Maximum == nil {
		maxima, err := dusk.GetObjectTransitMaximaTime(datetime, eq, latitude, longitude)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		transit.Maximum = maxima
	}

	// Create the Maximum gin.H JSON object representation:
	maximum := GetStandardTransitProperties(transit.Maximum, eq, longitude, latitude)

	// Create the Set gin.H JSON object representation:
	set := GetStandardTransitProperties(transit.Set, eq, longitude, latitude)

	path, _ := dusk.GetObjectHorizontalCoordinatesForDay(datetime, eq, longitude, latitude)

	c.JSON(http.StatusOK, gin.H{
		"observer": observer,
		"rise":     rise,
		"maximum":  maximum,
		"set":      set,
		"path":     path,
	})
}
