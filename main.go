/*
 * Main entry point for the web server
 */

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/daniel-vera-g/go-server-template/app"
	"github.com/daniel-vera-g/go-server-template/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	// Sign up & login
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	// Create, retrieve, change & delete note
	router.HandleFunc("/api/notes/new", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/api/me/notes", controllers.GetNotesFor).Methods("GET")
	// TODO
	// router.HandleFunc("/api/me/notes/change", controllers.ChangeNote).Methods("PUT")
	// TODO
	// router.HandleFunc("/api/me/notes/delete/" controllers.DeleteNote).Methods("DELETE")

	// router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Server running on Port: ", port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
