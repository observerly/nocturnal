package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	middleware "github.com/observerly/nocturnal/internal/middleware"
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

	// Setup Cross Origin Resource Sharing:
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

	// Setup Helmet Security Headers:
	r.Use(middleware.HelmetMiddleware())

	// Initialise Sentry if GIN_MODE is release and DSN is set:
	dsn := os.Getenv("SENTRY_DSN")

	// If we have a dsn, print out an obfuscated version of it with the last 10 characters replace with '*':
	if dsn != "" {
		fmt.Printf("SENTRY_DSN: %v\n", strings.Repeat("*", len(dsn)-10)+dsn[len(dsn)-10:])
	}

	// If we are in release mode and have a DSN, initialise Sentry:
	if mode == "release" && dsn != "" {
		// Make a log that we are initialising Sentry:
		fmt.Println("Initialising Sentry...")

		// To initialize Sentry's handler, you need to initialize Sentry itself beforehand:
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:           dsn,
			EnableTracing: true,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
		}); err != nil {
			fmt.Printf("Sentry initialization failed: %v\n", err)
		}

		// Use sentrygin middleware to send errors to Sentry:
		r.Use(sentrygin.New(sentrygin.Options{}))
	}

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
