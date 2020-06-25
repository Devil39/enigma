package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Devil39/enigma/api/handlers"
	questions "github.com/Devil39/enigma/pkg/question"
	"github.com/Devil39/enigma/pkg/user"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	mR := mux.NewRouter()
	authR := mR.PathPrefix("/api/auth").Subrouter()
	questionR := mR.PathPrefix("/api/question").Subrouter()

	db, err := sqlx.Connect("postgres", "user=postgres dbname=enigma password=1234")
	if err != nil {
		panic(err)
	}

	userRepo := user.NewPostgresRepo(db)
	quesRepo := questions.MakeNewQuestionRepo(db)

	userSvc := user.NewUserService(userRepo)
	quesSvc := questions.MakeNewQuestionService(quesRepo)

	handlers.MakeAuthHandler(authR, userSvc)
	handlers.MakeQuestionsHandler(questionR, quesSvc, userSvc)

	srv := http.Server{
		Handler:      mR,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}
