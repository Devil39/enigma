package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/Devil39/enigma/pkg/models"
	"github.com/Devil39/enigma/pkg/user"
	"github.com/gorilla/mux"
)

type loginRequest struct {
	emailID  string `json:"emailId"`
	password string `json:"password"`
}

func signupHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userSvc.CreateUser()
	}
}

func loginHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req loginRequest

		dec := json.NewDecoder(r.Body)
		dec.Decode(&req)

		fmt.Println(req)

		defer r.Body.Close()

		userSvc.Login()
	}
}

//MakeAuthHandler defines the route handlers for auth
func MakeAuthHandler(r *mux.Router, userSvc user.Service) {
	r.HandleFunc("/signup", signupHandler(userSvc)).Methods("POST")
	r.HandleFunc("/login", loginHandler(userSvc)).Methods("POST")
}
