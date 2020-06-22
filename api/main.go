package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Devil39/enigma/api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	mR := mux.NewRouter()
	authR := mR.PathPrefix("/api/auth").Subrouter()
	handlers.MakeAuthHandler(authR)

	srv := http.Server{
		Handler:      authR,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}
