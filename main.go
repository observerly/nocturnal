package main

import (
	"flag"
	"log"

	"github.com/observerly/nocturnal/internal/router"
	"github.com/observerly/nocturnal/pkg/moon"
)

var (
	port = flag.String("port", ":8103", "Port to listen on. Default is 8103.")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	r := router.SetupRouter()

	r.GET("/api/v1/moon", moon.GetMoon)

	// Listen on port
	log.Fatal(r.Run(*port))
}
