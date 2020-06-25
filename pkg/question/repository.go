package questions

import "github.com/Devil39/enigma/pkg/entities"

//Repository represents question repository
type Repository interface {
	AddQuestion(AddQuestionReq) error
	GetAllQuestions() ([]entities.Question, error)
	CheckAnswer(CheckAnswerReq) (bool, error)
	GetHint(string) (string, error)
}

//AddQuestionReq represents the request for add a new question
type AddQuestionReq struct {
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Answer string `json:"answer"`
}

//CheckAnswerReq represents the request for check answer
type CheckAnswerReq struct {
	ID     string `json:"id"`
	Answer string `json:"answer"`
}
