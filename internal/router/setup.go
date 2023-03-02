package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getAPIVersionFromEnv() string {
	value, exists := os.LookupEnv("API_VERSION_LATEST")
	if !exists {
		value = "v2"
	}
	return value
}

func CORSMiddleware(config cors.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if config.AllowAllOrigins || config.AllowOrigins != nil {
			for _, value := range config.AllowOrigins {
				if value == origin {
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				}
			}
		}

		if config.AllowMethods != nil {
			c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ","))
		}

		if config.AllowHeaders != nil {
			c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ","))
		}

		if config.ExposeHeaders != nil {
			c.Writer.Header().Set("Access-Control-Expose-Headers", strings.Join(config.ExposeHeaders, ","))
		}

		if config.AllowCredentials {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if config.MaxAge > 0 {
			c.Writer.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", config.MaxAge))
		}

		c.Next()
	}
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

	config := cors.Config{
		AllowOrigins: []string{
			"https://observerly.com",
			"https://app.observerly.com",
			"https://vega.observerly.com",
			"http://localhost:3001",
		},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}

	r.Use(CORSMiddleware(config))

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
				"endpoint":    "/api/v1",
				"name":        "Nocturnal API by observerly",
			},
		)
	})

	r.GET("/api/v2", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"description": "Nocturnal 🌑 is observerly's Gin Gonic API for Lunar, Solar and astronomical advanced scheduling, that utilises Dusk.",
				"endpoint":    fmt.Sprintf("/api/%v", version),
				"name":        "Nocturnal API by observerly",
			},
		)
	})

	// 404 Handler, ensure we are always redirected from api to the latest version of the API:
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/api/v2")
	})

	return r
}
