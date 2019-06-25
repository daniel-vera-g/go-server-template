package main

import (
	"fmt"
	"log"
	"net/http"
)

// Function of the type http.HandleFunc
// Takes a ResponseWriter to assemble the HTTP response
// and a Request with the client HTTP request data
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %s", r.URL.Path[1:])
}

func main() {
	// Handle all root request with the handler function
	http.HandleFunc("/", handler)
	// Listen on Port 8080 and return(log) when there is an unexpected error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
