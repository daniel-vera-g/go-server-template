package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccount(t *testing.T) {

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/user/new", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf(""))
	Handler := http.HandleFunc(CreateAccount)
	Handler.ServeHTTP(rr, req)

}
