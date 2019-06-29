/*
 * Main entry point for the web server
 */

package main

import (
	"fmt"
	"net/http"
	"os"

	"./app"
	"./controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Sign up & login
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	// Create, retrieve, change & delete note
	router.HandleFunc("/api/notes/new", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/api/me/notes", controllers.GetNoteFor).Methods("GET") //  user/2/notes
	router.HandleFunc("/api/me/notes/change", controllers.ChangeNote).Methods("PUT")
	router.HandleFunc("/api/me/notes/delete/" controllers.DeleteNote).Methods("DELETE")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Server running on Port: %s", port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
