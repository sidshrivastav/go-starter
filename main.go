package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-starter/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there %s", r.URL.Path[1:])
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify the environment (dev/prod)")
		return
	}
	env := os.Args[1]
	cfg := config.LoadConfig(env)

	http.HandleFunc("/", handler)
	// Start HTTP server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
