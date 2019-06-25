package main

import (
	"fmt"
	"io/ioutil"
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

func main() {

	// Create, save and load new page
	p1 := &Page{Title: "TestPage", Body: []byte("A simple message")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
