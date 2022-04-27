package main

import (
	"flag"
	"log"

	"github.com/observerly/nocturnal/internal/router"
	"github.com/observerly/nocturnal/pkg/moon"
	"github.com/observerly/nocturnal/pkg/sun"
	"github.com/observerly/nocturnal/pkg/twilight"
)

var (
	port = flag.String("port", ":8103", "Port to listen on. Default is 8103.")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	r := router.SetupRouter()

	// Moon Properties API
	r.GET("/api/v1/moon", moon.GetMoon)

	// Sun Properties API
	r.GET("/api/v1/sun", sun.GetSun)

	// Twilight Properties API
	r.GET("/api/v1/twilight", twilight.GetTwilight)

	// Listen on port
	log.Fatal(r.Run(*port))
}
