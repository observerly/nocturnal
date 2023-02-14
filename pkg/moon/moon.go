package moon

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/observerly/dusk/pkg/dusk"
	"github.com/observerly/nocturnal/internal/query"
	"github.com/observerly/nocturnal/internal/utils"
)

// GET /moon
func GetMoonDeprecatedV1(c *gin.Context) {
	d, lon, lat := query.GetDefaultObserverParams(c)

	datetime, _ := utils.ParseDatetimeRFC3339(d)

	longitude, _ := strconv.ParseFloat(lon, 64)

	latitude, _ := strconv.ParseFloat(lat, 64)

	ec := dusk.GetLunarEclipticPosition(datetime)

	eq := dusk.GetLunarEquatorialPosition(datetime)

	hz := dusk.ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, eq)

	ph := dusk.GetLunarPhase(datetime, longitude, ec)

	rs, _ := dusk.GetMoonriseMoonsetTimes(datetime, longitude, latitude)

	observer := gin.H{
		"datetime":  datetime,
		"longitude": longitude,
		"latitude":  latitude,
	}

	position := gin.H{
		"alt": hz.Altitude,
		"az":  hz.Azimuth,
		"ra":  eq.RightAscension,
		"dec": eq.Declination,
	}

	phase := gin.H{
		"age":          ph.Days,
		"angle":        ph.Angle,
		"d":            ph.Age,
		"fraction":     ph.Fraction,
		"illumination": ph.Illumination,
	}

	transit := gin.H{}

	if rs.Rise.IsZero() {
		transit["rise"] = nil
	} else {
		transit["rise"] = rs.Rise.Format(time.RFC3339)
	}

	if rs.Set.IsZero() {
		transit["set"] = nil
	} else {
		transit["set"] = rs.Set.Format(time.RFC3339)
	}

	c.JSON(http.StatusOK, gin.H{
		"observer": observer,
		"position": position,
		"phase":    phase,
		"transit":  transit,
	})
}

func GetStandardLunarProperties(datetime time.Time, longitude float64, latitude float64) gin.H {
	ec := dusk.GetLunarEclipticPositionLawrence(datetime.UTC())

	eq := dusk.ConvertEclipticCoordinateToEquatorial(datetime.UTC(), ec)

	hz := dusk.ConvertEquatorialCoordinateToHorizontal(datetime.UTC(), longitude, latitude, eq)

	ph := dusk.GetLunarPhase(datetime.UTC(), longitude, ec)

	airmass := dusk.GetRelativeAirMass(hz.Altitude)

	refraction := dusk.GetAtmosphericRefraction(hz.Altitude)

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
	}
}

// GET /moon v2
func GetMoon(c *gin.Context) {
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

	// Get the next Moon rise and set times:
	rs, _ := dusk.GetMoonriseMoonsetTimes(datetime, longitude, latitude)

	// Calculate Lunar properties (e.g., phase) at the datetime of the next rise:
	var rise gin.H = nil

	if !rs.Rise.IsZero() {
		rise = GetStandardLunarProperties(rs.Rise, longitude, latitude)
	}

	// Calculate Lunar properties (e.g., phase) at the datetime of the next set:
	var set gin.H = nil

	if !rs.Set.IsZero() {
		set = GetStandardLunarProperties(rs.Set, longitude, latitude)
	}

	c.JSON(http.StatusOK, gin.H{
		"observer": observer,
		"rise":     rise,
		"set":      set,
	})
}
