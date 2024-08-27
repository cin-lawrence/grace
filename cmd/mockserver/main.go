package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register the handler function for the root path "/"
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8123")
	if err := http.ListenAndServe(":8123", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
