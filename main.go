package main

import (
	"flag"
	"log"

	"github.com/observerly/nocturnal/internal/router"
	"github.com/observerly/nocturnal/pkg/moon"
	"github.com/observerly/nocturnal/pkg/sun"
	"github.com/observerly/nocturnal/pkg/transit"
	"github.com/observerly/nocturnal/pkg/twilight"
)

var (
	port = flag.String("port", ":8103", "Port to listen on. Default is 8103.")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	r := router.SetupRouter()

	// Moon (Lunar) Properties API
	r.GET("/api/v1/moon", moon.GetMoon)
	r.GET("/api/v1/lunar", moon.GetMoon)

	// Sun (Solar) Properties API
	r.GET("/api/v1/sun", sun.GetSun)
	r.GET("/api/v1/solar", sun.GetSun)

	// Transit Properties API
	r.GET("/api/v1/transit", transit.GetTransit)

	// Twilight (Crepusculum) Properties API
	r.GET("/api/v1/twilight", twilight.GetTwilight)

	// Listen on port
	log.Fatal(r.Run(*port))
}
