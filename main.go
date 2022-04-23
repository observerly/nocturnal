package main

import (
	"flag"
	"log"

	"github.com/observerly/nocturnal/internal/router"
)

var (
	port = flag.String("port", ":8103", "Port to listen on. Default is 8103.")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	r := router.SetupRouter()

	// Listen on port
	log.Fatal(r.Run(*port))
}
