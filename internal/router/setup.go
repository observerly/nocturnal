package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getAPIVersionFromEnv() string {
	value, exists := os.LookupEnv("API_VERSION_LATEST")
	if !exists {
		value = "v1"
	}
	return value
}

func SetupRouter() *gin.Engine {
	var version = getAPIVersionFromEnv()

	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create gin router
	r := gin.Default()

	// Logging middleware
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err,
				},
			)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://observerly.com", "https://app.observerly.com"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/version", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"latest": fmt.Sprintf("/api/%v", version),
			},
		)
	})

	// /api/v1 Group w/ Name(v1)
	r.GET("/api/v1", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"description": "Nocturnal 🌑 is observerly's Gin Gonic API for Lunar and Solar advanced scheduling, that utilises Dusk.",
				"endpoint":    fmt.Sprintf("/api/%v", version),
				"name":        "Nocturnal API by observerly",
			},
		)
	})

	// 404 Handler, ensure we are always redirected from api to the latest version of the API:
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/api/v1")
	})

	return r
}
