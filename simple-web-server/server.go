package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Function of the type http.HandleFunc
// Takes a ResponseWriter to assemble the HTTP response
// and a Request with the client HTTP request data
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %s", r.URL.Path[1:])
}

/**
* TODO export to external file for separation of concerns
**/

// Data stricture representing the page title & body
type Page struct {
	Title string
	// Byte slice to store the body(Type expected by io)
	Body []byte
}

// Method to save a text file to the system
// Takes as its receiver p, a pointer to the Page struct
// Returns error or nil
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// 0600 = Read + write permissions
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Function to load pages
// Takes The title of the page as parameter
// Returns a Pointer to a Page with the Title and body
// or an error if the file could not be read
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

// View handler to view a page
// Handle URLs with the prefix /view/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// Get the page to be shown
	title := r.URL.Path[len("/view/"):]
	// Load the page data
	// TODO error handling
	p, _ := loadPage(title)
	// Format the page with HTML and write it to w(ResponseWriter)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	// Handle all root request with the handler function
	// http.HandleFunc("/", handler)

	// Handle requests for a specific site
	http.HandleFunc("/view/", viewHandler)
	// Server the page on Port 8080 and return if there is an unexpected error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
