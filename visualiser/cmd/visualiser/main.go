package main

import (
	"flag"
	"log"
	"time"

	"github.com/allantabilog/visualiser/examples"
	"github.com/allantabilog/visualiser/internal/server"
	"github.com/allantabilog/visualiser/internal/visualizer"
)

func main() {
	// Parse command line flags
	runDemo := flag.Bool("demo", false, "Run demonstration algorithms")
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	// Create a new visualizer
	viz := visualizer.NewVisualizer()
	
	// Create and start the server
	srv := server.NewServer(viz)
	
	log.Printf("Data Structure Visualizer starting...")
	log.Printf("Open your browser to http://localhost:%d", *port)
	
	// Optionally run demo algorithms
	if *runDemo {
		log.Printf("Starting demonstration algorithms in 3 seconds...")
		go func() {
			time.Sleep(3 * time.Second)
			examples.RunAllDemos(viz)
		}()
	}
	
	if err := srv.Start(*port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}