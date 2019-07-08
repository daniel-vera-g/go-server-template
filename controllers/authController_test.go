package controllers

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// Suite for the test environement
type HandlerSuite struct {
	suite.Suite
	// Auth
	uname string
	psswd string
	// Notes
	jwtToken string
	name     string
	note     string
}

func (suite *HandlerSuite) SetupSuite() {
	// Generate random string for email & password
	rand.Seed(time.Now().UnixNano())
	psswd := RandSeq(10)
	uname := RandSeq(8)

	// Set up credentials
	credentials := LoginCredentials{
		Email:    uname + "@test.com",
		Password: psswd,
	}

	// generate json
	jsonPayload, err := json.Marshal(credentials)

	// HTTP request ressource
	ressource := "/api/user/new"

	// Request object
	req, err := http.NewRequest("POST", ressource, bytes.NewBuffer(jsonPayload))
	if err != nil {
		suite.T().Fatal(err)
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")

	// Record & make request
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)
	handler.ServeHTTP(res, req)

	// Get json response
	// var resJson map[string]interface{}
	var jwtRes jwtResponse
	if err := json.NewDecoder(res.Body).Decode(&jwtRes); err != nil {
		panic(err)
	}

	// Token for further Testing
	suite.jwtToken = jwtRes.Account.Token
	suite.name = RandSeq(8)
	suite.note = RandSeq(8)

	// Generate random string for email & password
	rand.Seed(time.Now().UnixNano())
	suite.uname = RandSeq(8)
	suite.psswd = RandSeq(10)
}

// Activate the test suite
func TestSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

// Credentials for account testing
type LoginCresentials struct {
	Email    string
	Password string
}

var psswd, uname string

func (suite *HandlerSuite) TestCreateAccount() {

	// Set up credentials
	credentials := LoginCresentials{
		Email:    suite.uname + "@test.com",
		Password: suite.psswd,
	}

	// generate json
	jsonPayload, err := json.Marshal(credentials)

	// HTTP request ressource
	ressource := "/api/user/new"

	// Request object
	req, err := http.NewRequest("POST", ressource, bytes.NewBuffer(jsonPayload))
	if err != nil {
		suite.T().Fatal(err)
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")

	// Record & make request
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)
	handler.ServeHTTP(res, req)

	// Get json response
	var resJson map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resJson); err != nil {
		panic(err)
	}

	// Actual test
	if resJson["message"] != "Account has been created" && resJson["status"] != "true" {
		suite.T().Errorf("Handler returned wrong status. Got %s want true", resJson["status"])
	}
}

func (suite *HandlerSuite) TestAuthenticate() {

	// Set up credentials
	// credentials := LoginCresentials{
	// Email:    suite.uname + "@test.com",
	// Password: suite.psswd,
	// }
	credentials := LoginCresentials{
		Email:    "test@test.com",
		Password: "password",
	}

	// generate json
	jsonPayload, err := json.Marshal(credentials)

	// TODO refactor(First make account)

	// Request object
	reqNew, err := http.NewRequest("POST", "/api/new/user", bytes.NewBuffer(jsonPayload))
	if err != nil {
		suite.T().Fatal(err)
	}

	// Set header
	reqNew.Header.Set("Content-Type", "application/json")

	// Record & make request
	resNew := httptest.NewRecorder()
	handlerNew := http.HandlerFunc(CreateAccount)
	handlerNew.ServeHTTP(resNew, reqNew)

	//TODO refactor

	// HTTP request ressource
	ressource := "/api/user/login"

	// Request object
	req, err := http.NewRequest("POST", ressource, bytes.NewBuffer(jsonPayload))
	if err != nil {
		suite.T().Fatal(err)
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")

	// Record & make request
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(Authenticate)
	handler.ServeHTTP(res, req)

	// Get json response
	var resJson map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resJson); err != nil {
		panic(err)
	}

	// Actual test
	if resJson["message"] != "Logged In" && resJson["status"] != "true" {
		suite.T().Errorf("Handler returned wrong status. Got %v want true", resJson["status"])
		suite.T().Errorf("The message was: %s", resJson["message"])
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Create random strings with specific length
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Clear up databases after test TODO
