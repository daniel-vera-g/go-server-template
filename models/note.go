/*
 * Let users create/store notes
 */

package models

import (
	"fmt"
	u "utils"
)

type Note struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	UserId uint   `json:"user_id"` //The user that this note belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (note *note) Validate() (map[string]interface{}, bool) {

	if note.Name == "" {
		return u.Message(false, "note name should be on the payload"), false
	}

	if note.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if note.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (note *note) Create() map[string]interface{} {

	if resp, ok := note.Validate(); !ok {
		return resp
	}

	// Add Note
	insert, err := db.Query("INSERT INTO notes VALUES ()")

	GetDB().Create(note)

	resp := u.Message(true, "success")
	resp["note"] = note
	return resp
}

func Getnote(id uint) *note {

	note := &note{}
	err := GetDB().Table("notes").Where("id = ?", id).First(note).Error
	if err != nil {
		return nil
	}
	return note
}

func Getnotes(user uint) []*note {

	notes := make([]*note, 0)
	err := GetDB().Table("notes").Where("user_id = ?", user).Find(&notes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return notes
}
