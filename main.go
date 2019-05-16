package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Setup HTTP handlers
	http.HandleFunc("/", handle)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %v", r.URL.Path)
	fmt.Fprint(w, "Hello from Go")
}
