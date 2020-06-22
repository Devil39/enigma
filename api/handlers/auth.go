package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Namastey Duniyaa!")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Namastey Duniyaa!")
}

//MakeAuthHandler defines the route handlers for auth
func MakeAuthHandler(r *mux.Router) {
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/signup", signupHandler).Methods("POST")
}
