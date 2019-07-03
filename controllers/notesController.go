/*
 * Create and get notes
 */
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/daniel-vera-g/go-server-template/models"
	u "github.com/daniel-vera-g/go-server-template/utils"
)

var CreateNote = func(w http.ResponseWriter, r *http.Request) {

	//Grab the id of the user that send the request
	user := r.Context().Value("user").(uint)
	note := &models.Note{}

	// Get the note from the request
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	// Create note & send response
	note.UserId = user
	resp := note.Create()
	u.Respond(w, resp)
}

var GetNotesFor = func(w http.ResponseWriter, r *http.Request) {

	// Get id of user + notes of user, & send notes
	id := r.Context().Value("user").(uint)
	data := models.GetNotes(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
