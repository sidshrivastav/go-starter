package main

import (
	"fmt"
	"log"
	"net/http"

	"go-starter/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there %s", r.URL.Path[1:])
}

func main() {
	// Load the application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s\n", err)
	}

	// Print loaded configurations
	fmt.Printf("App Name: %s\n", cfg.AppName)
	fmt.Printf("App Version: %s\n", cfg.AppVersion)

	// Database configuration
	fmt.Printf("Connecting to DB: %s:%d with user %s\n", cfg.Database.Host, cfg.Database.Port, cfg.Database.User)

	// Logging configuration
	fmt.Printf("Logging level: %s\n", cfg.Logging.Level)

	http.HandleFunc("/", handler)
	// Start HTTP server
	addr := fmt.Sprintf(":%d", cfg.AppPort)
	log.Fatal(http.ListenAndServe(addr, nil))
}
