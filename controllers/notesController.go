/*
 * Create and get notes
 */
package controllers

import (
	"encoding/json"
	"net/http"

	u "./utils"
	"github.com/daniel-vera-g/go-server-template/models"
)

var CreateNote = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	note := &models.Note{}

	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	note.UserId = user
	resp := note.Create()
	u.Respond(w, resp)
}

var GetNotesFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetNotes(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
