/*
 * Let users create/store notes
 */
package models

import (
	"fmt"

	u "../utils"
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Name   string `json:"name"`
	Note   string `json:"note"`
	UserId uint   `json:"user_id"` //The user that this note belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (note *Note) Validate() (map[string]interface{}, bool) {

	if note.Name == "" {
		return u.Message(false, "Note name should be on the payload"), false
	}

	if note.Note == "" {
		return u.Message(false, "Note text should be on the payload"), false
	}

	if note.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (note *Note) Create() map[string]interface{} {

	if resp, ok := note.Validate(); !ok {
		return resp
	}

	GetDB().Create(note)

	resp := u.Message(true, "success")
	resp["note"] = note
	return resp
}

func GetNote(id uint) *Note {

	note := &Note{}
	err := GetDB().Table("notes").Where("id = ?", id).First(note).Error
	if err != nil {
		return nil
	}
	return note
}

func GetNotes(user uint) []*Note {

	notes := make([]*Note, 0)
	err := GetDB().Table("notes").Where("user_id = ?", user).Find(&notes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return notes
}
