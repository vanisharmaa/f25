package main

// Import necessary packages
import (
	"fmt"      // For printing/logging and sending formatted text
	"net/http" // For building HTTP servers and handling requests
)

func main() {
	// Register a route "/health" with a handler function
	// HandleFunc takes: path string, and handler func(w, r)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Write "Auth Service OK" to the HTTP response
		// fmt.Fprintln writes to any io.Writer (here it's the HTTP response)
		fmt.Fprintln(w, "Auth Service OK")
	})

	// Log to the console that the service is running
	// This is just for your own reference in the terminal
	fmt.Println("Auth service running on :8080")

	// Start the HTTP server on port 8080
	// ":8080" means listen on all network interfaces on port 8080
	// nil means use the default multiplexer (where we registered /health)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// If the server fails to start or crashes, log the error
		fmt.Println("Server failed:", err)
	}
}
