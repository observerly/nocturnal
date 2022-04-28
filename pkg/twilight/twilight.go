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

	observer := gin.H{
		"datetime":  datetime,
		"longitude": fmt.Sprintf("%f", longitude),
		"latitude":  fmt.Sprintf("%f", latitude),
	}

	// Civil Twilight:

	civil, location, _ := dusk.GetLocalCivilTwilight(datetime, longitude, latitude, 0)

	ct := gin.H{
		"from":     civil.From.Format(time.RFC3339),
		"until":    civil.Until.Format(time.RFC3339),
		"duration": float64(civil.Duration.Milliseconds()) * 0.001 / 3600,
		"location": location.String(),
		"horizon":  -6,
	}

	// Nautical Twilight:

	nautical, location, _ := dusk.GetLocalNauticalTwilight(datetime, longitude, latitude, 0)

	nt := gin.H{
		"from":     nautical.From.Format(time.RFC3339),
		"until":    nautical.Until.Format(time.RFC3339),
		"duration": float64(nautical.Duration.Milliseconds()) * 0.001 / 3600,
		"location": location.String(),
		"horizon":  -12,
	}

	// Astronomical Twilight:

	astronomical, location, _ := dusk.GetLocalAstronomicalTwilight(datetime, longitude, latitude, 0)

	at := gin.H{
		"from":     astronomical.From.Format(time.RFC3339),
		"until":    astronomical.Until.Format(time.RFC3339),
		"duration": float64(astronomical.Duration.Milliseconds()) * 0.001 / 3600,
		"location": location.String(),
		"horizon":  -18,
	}

	c.JSON(http.StatusOK, gin.H{
		"observer":     observer,
		"astronomical": at,
		"civil":        ct,
		"nautical":     nt,
	})
}
