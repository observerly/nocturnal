package sun

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/observerly/dusk/pkg/dusk"
	"github.com/observerly/nocturnal/internal/utils"
)

// GET Sun
func GetSun(c *gin.Context) {
	d := c.DefaultQuery("datetime", time.Now().String())

	lon := c.DefaultQuery("longitude", strconv.Itoa(0))

	lat := c.DefaultQuery("latitude", strconv.Itoa(0))

	datetime, _ := utils.ParseDatetimeRFC3339(d)

	longitude, _ := strconv.ParseFloat(lon, 64)

	latitude, _ := strconv.ParseFloat(lat, 64)

	eq := dusk.GetSolarEquatorialPosition(datetime)

	hz := dusk.ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, eq)

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

	c.JSON(http.StatusOK, gin.H{
		"observer": observer,
		"position": position,
	})
}
