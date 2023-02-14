package query

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDefaultObserverParams(c *gin.Context) (string, string, string) {
	datetime := c.DefaultQuery("datetime", time.Now().Format(time.RFC3339))

	longitude := c.DefaultQuery("longitude", strconv.Itoa(0))

	latitude := c.DefaultQuery("latitude", strconv.Itoa(0))

	return datetime, longitude, latitude
}
