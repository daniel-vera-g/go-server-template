/*
 * Functions to build JSON messages
 * & return responses
 */

package utils

import (
	"encoding/json"
	"net/http"
)

// Build the JSON message with a status and a message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Attach the JSON to the response
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
