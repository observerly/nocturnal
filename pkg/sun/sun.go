package sun

import (
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

	rstoday, _ := dusk.GetSunriseSunsetTimes(datetime, 0, longitude, latitude, 0)

	rstomorrow, _ := dusk.GetSunriseSunsetTimes(datetime.Add(time.Hour*24), 0, longitude, latitude, 0)

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

	transit := gin.H{
		"rise": rstoday.Rise.Format(time.RFC3339),
		"set":  rstoday.Set.Format(time.RFC3339),
	}

	tomorrow := gin.H{
		"rise": rstomorrow.Rise.Format(time.RFC3339),
		"set":  rstomorrow.Set.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, gin.H{
		"observer": observer,
		"position": position,
		"transit":  transit,
		"tomorrow": tomorrow,
	})
}
