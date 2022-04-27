package twilight

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/observerly/dusk/pkg/dusk"
	"github.com/observerly/nocturnal/internal/utils"
)

func GetTwilight(c *gin.Context) {
	d := c.DefaultQuery("datetime", time.Now().Format(time.RFC3339))

	lon := c.DefaultQuery("longitude", strconv.Itoa(0))

	lat := c.DefaultQuery("latitude", strconv.Itoa(0))

	datetime, _ := utils.ParseDatetimeRFC3339(d)

	longitude, _ := strconv.ParseFloat(lon, 64)

	latitude, _ := strconv.ParseFloat(lat, 64)

	t, location, _ := dusk.GetLocalAstronomicalTwilight(datetime, longitude, latitude, 0)

	observer := gin.H{
		"datetime":  datetime,
		"longitude": fmt.Sprintf("%f", longitude),
		"latitude":  fmt.Sprintf("%f", latitude),
	}

	twilight := gin.H{
		"from":     t.From.Format(time.RFC3339),
		"until":    t.Until.Format(time.RFC3339),
		"duration": float64(t.Duration.Milliseconds()) * 0.001 / 3600,
		"location": location.String(),
	}

	c.JSON(http.StatusOK, gin.H{
		"observer": observer,
		"twilight": twilight,
	})
}
