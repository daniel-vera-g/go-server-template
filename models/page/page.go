package page

// Data stricture representing the page title & body
type Page struct {
	Title string
	// Byte slice to store the body(Type expected by io)
	Body []byte
}
