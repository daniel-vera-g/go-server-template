package main

import (
	// TODO rm fmt
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
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
	renderTemplate(w, "view", p)
}

// editHandler loads the page
// creates an empty page struct if not existent
// displays an HTML form
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	// If not existent, create empty Page struct
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// renderTemplate renders a specific HTML page
// It takes the ResponseWriter, a name of the template and a pointer to the page
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// Read contents of HTML file & return contents
	t, _ := template.ParseFiles(tmpl + ".html")
	// Write generated HTML to the reponse
	t.Execute(w, p)
}

func main() {
	// Handle all root request with the handler function
	// http.HandleFunc("/", handler)

	// Reques handler
	// Specific site
	http.HandleFunc("/view/", viewHandler)
	// Edit page
	http.HandleFunc("/edit/", editHandler)
	// Save data
	// http.HandleFunc("/save/", saveHandler)

	// Server the page on Port 8080 and return if there is an unexpected error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
