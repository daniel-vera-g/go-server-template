package main

import (
	"log"
	"net/http"

	handlers "./routes"
)

func main() {
	// Request handler
	// Specific site
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	// Edit page
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	// Save data
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))

	// Server the page on Port 8080 and return if there is an unexpected error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
