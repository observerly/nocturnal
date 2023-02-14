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
func GetMoon(c *gin.Context) {
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
