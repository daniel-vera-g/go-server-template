package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/daniel-vera-g/go-server-template/app"
)

// Credentials for account testing
type LoginCredentials struct {
	Email    string
	Password string
}

type jwtResponse struct {
	Account account
}

type account struct {
	Token string
}

type noteRequest struct {
	Name string
	Note string
}

func (suite *HandlerSuite) TestCreateNote() {

	// Note informations
	credentials := noteRequest{
		Name: suite.name,
		Note: suite.note,
	}

	// generate json
	jsonPayload, err := json.Marshal(credentials)

	// Request object
	req, err := http.NewRequest("POST", "/api/notes/new", bytes.NewBuffer(jsonPayload))
	if err != nil {
		suite.T().Fatal(err)
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+suite.jwtToken)

	// Record & make request
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateNote)
	handlerToTest := app.JwtAuthentication(handler)
	handlerToTest.ServeHTTP(res, req)

	// Get json response
	var resJson map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resJson); err != nil {
		panic(err)
	}

	fmt.Println("The response:", resJson)

	// Actual test
	if resJson["message"] != "success" && resJson["status"] != "true" {
		suite.T().Errorf("Handler returned wrong status. Got %v want true", resJson["status"])
		suite.T().Errorf("The message was: %s", resJson["message"])
	}

}
