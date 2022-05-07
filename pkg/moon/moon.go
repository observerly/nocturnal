package moon

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/observerly/dusk/pkg/dusk"
	"github.com/observerly/nocturnal/internal/utils"
)

// GET /moon
func GetMoon(c *gin.Context) {
	d := c.DefaultQuery("datetime", time.Now().String())

	lon := c.DefaultQuery("longitude", strconv.Itoa(0))

	lat := c.DefaultQuery("latitude", strconv.Itoa(0))

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
		"longitude": fmt.Sprintf("%f", longitude),
		"latitude":  fmt.Sprintf("%f", latitude),
	}

	position := gin.H{
		"alt": fmt.Sprintf("%f", hz.Altitude),
		"az":  fmt.Sprintf("%f", hz.Azimuth),
		"ra":  fmt.Sprintf("%f", eq.RightAscension),
		"dec": fmt.Sprintf("%f", eq.Declination),
	}

	phase := gin.H{
		"age":          fmt.Sprintf("%f", ph.Days),
		"angle":        fmt.Sprintf("%f", ph.Angle),
		"d":            fmt.Sprintf("%f", ph.Age),
		"fraction":     fmt.Sprintf("%f", ph.Fraction),
		"illumination": fmt.Sprintf("%f", ph.Illumination),
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
