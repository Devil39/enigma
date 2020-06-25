package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Devil39/enigma/api/views"
	"github.com/Devil39/enigma/pkg"
	"github.com/Devil39/enigma/pkg/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func signupHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := make(map[string]interface{})

		req := user.AuthRequest{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&req)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, pkg.ErrWrongFormat.Error(), message)
			return
		}

		defer r.Body.Close()

		user, err := userSvc.CreateUser(req)
		if err == pkg.ErrInvalidEmailAndPass {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusUnauthorized, err.Error(), message)
			return
		} else if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		token, err := createToken(user.UUID, user.EmailID)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully signed up!"
		message["token"] = token
		views.SendResponse(w, http.StatusCreated, "", message)
	}
}

func loginHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req user.AuthRequest
		message := make(map[string]interface{})

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&req)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, pkg.ErrWrongFormat.Error(), message)
			return
		}

		defer r.Body.Close()

		user, err := userSvc.Login(req)
		if err == pkg.ErrInvalidEmailAndPass {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusUnauthorized, err.Error(), message)
			return
		} else if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		token, err := createToken(user.UUID, user.EmailID)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully logged in!"
		message["token"] = token
		views.SendResponse(w, http.StatusOK, "", message)
	}
}

//MakeAuthHandler defines the route handlers for auth
func MakeAuthHandler(r *mux.Router, userSvc user.Service) {
	r.HandleFunc("/signup", (signupHandler(userSvc))).Methods("POST")
	r.HandleFunc("/login", (loginHandler(userSvc))).Methods("POST")
}

func createToken(userid, emailid string) (string, error) {

	//var token Token
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email_id"] = emailid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte("ENIGMA"))
	if err != nil {
		return "", err
	}

	return token, nil
}
