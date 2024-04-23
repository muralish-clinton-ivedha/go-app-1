package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "Hello, World!" print statement
	fmt.Println("Hello, World!")

	// Define HTTP routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api1", api1Handler)
	http.HandleFunc("/api2", api2Handler)
	http.HandleFunc("/api3", api3Handler)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// Handler for the root endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Handler for /api1 endpoint
func api1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is API 1")
}

// Handler for /api2 endpoint
func api2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is API 2")
}

// Handler for /api3 endpoint
func api3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is API 3")
}
