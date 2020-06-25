package handlers

import (
	"encoding/json"
	"net/http"

	middlewares "github.com/Devil39/enigma/api/middleware"
	"github.com/Devil39/enigma/api/views"
	"github.com/Devil39/enigma/pkg"
	questions "github.com/Devil39/enigma/pkg/question"
	"github.com/Devil39/enigma/pkg/user"

	"github.com/gorilla/mux"
)

func addQuestionHandler(quesSvc questions.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req questions.AddQuestionReq
		message := make(map[string]interface{})

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&req)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, pkg.ErrWrongFormat.Error(), message)
			return
		}
		defer r.Body.Close()

		if req.Answer == "" || req.Desc == "" || req.Title == "" {
			message["message"] = "Please check all the fields are present, and not empty"
			views.SendResponse(w, http.StatusBadRequest, pkg.ErrWrongFormat.Error(), message)
			return
		}

		err = quesSvc.AddQuestion(req)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully added question"
		views.SendResponse(w, http.StatusOK, "", message)
	}
}

func getAllQuestionsHandler(quesSvc questions.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := make(map[string]interface{})
		questions, err := quesSvc.GetAllQuestions()
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully retrieved questions!"
		message["questions"] = questions
		views.SendResponse(w, http.StatusOK, "", message)
	}
}

func checkAnswerHandler(quesSvc questions.Service, userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req questions.CheckAnswerReq
		message := make(map[string]interface{})

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&req)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, pkg.ErrWrongFormat.Error(), message)
			return
		}
		defer r.Body.Close()

		ok, err := quesSvc.CheckAnswer(req)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, err.Error(), message)
			return
		}

		if ok {
			message["message"] = "Correct Answer!"
		} else {
			message["message"] = "Incorrect Answer!"
			views.SendResponse(w, http.StatusOK, "", message)
			return
		}

		emailID := getEmailIDFromToken(r)

		err = userSvc.AddSolvedQuestion(emailID, req.ID)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["valid"] = ok
		views.SendResponse(w, http.StatusOK, "", message)
	}
}

func getHintHandler(quesSvc questions.Service, userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req questions.CheckAnswerReq
		message := make(map[string]interface{})

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&req)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, pkg.ErrWrongFormat.Error(), message)
			return
		}
		defer r.Body.Close()

		hint, err := quesSvc.GetHint(req.ID)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, err.Error(), message)
			return
		}

		emailID := getEmailIDFromToken(r)

		err = userSvc.AddHintUsed(emailID, req.ID)
		if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["hint"] = hint
		views.SendResponse(w, http.StatusOK, "", message)
	}
}

//MakeQuestionsHandler defines all the route for questions
func MakeQuestionsHandler(r *mux.Router, quesSvc questions.Service, userSvc user.Service) {
	r.HandleFunc("/add-question", middlewares.JwtMiddleware(addQuestionHandler(quesSvc))).Methods("POST")
	r.HandleFunc("/get-all-questions", middlewares.JwtMiddleware(getAllQuestionsHandler(quesSvc))).Methods("GET")
	r.HandleFunc("/check-answer", middlewares.JwtMiddleware(checkAnswerHandler(quesSvc, userSvc))).Methods("POST")
	r.HandleFunc("/get-hint", middlewares.JwtMiddleware(getHintHandler(quesSvc, userSvc))).Methods("POST")
}

func getEmailIDFromToken(r *http.Request) string {
	_, claims, _ := middlewares.VerifyToken(r)
	if claimStr, ok := claims["email_id"].(string); ok {
		return claimStr
	}
	return ""
}
