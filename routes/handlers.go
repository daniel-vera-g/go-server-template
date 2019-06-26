package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"text/template"
)

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

// Parse files once at Programm initialization
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// Store validation expression -> Only specific pages
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// getTitle extracts the page title with the validPath Global variable
// It returns a 404 error or the title of the page
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}

// renderTemplate renders a specific HTML page
// It takes the ResponseWriter, a name of the template and a pointer to the page
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// View handler to view a page
// Handle URLs with the prefix /view/
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// Load the page data
	p, err := loadPage(title)
	// If the page does not exist redirect them to edit
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// editHandler loads the page
// creates an empty page struct if not existent
// displays an HTML form
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	// If not existent, create empty Page struct
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// Handle submission of forms from the edit page
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// Load the page data
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Wrapper funtion to take a handler function & return a http.ResponseWriter function
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the page title from request + do validation
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// Call the provided handler `fn`
		fn(w, r, m[2])
	}
}
