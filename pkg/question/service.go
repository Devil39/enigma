package questions

import (
	"github.com/Devil39/enigma/pkg/entities"
)

//Service represents question usecases
type Service interface {
	AddQuestion(AddQuestionReq) error
	GetAllQuestions() ([]entities.Question, error)
	CheckAnswer(CheckAnswerReq) (bool, error)
	GetHint(string) (string, error)
}

type quesSvc struct {
	r Repository
}

//MakeNewQuestionService takes an instance of Questoin Repository and returns an instnace of question service
func MakeNewQuestionService(r Repository) Service {
	return &quesSvc{r: r}
}

func (q *quesSvc) AddQuestion(req AddQuestionReq) error {
	return q.r.AddQuestion(req)
}

func (q *quesSvc) GetAllQuestions() ([]entities.Question, error) {
	return q.r.GetAllQuestions()
}

func (q *quesSvc) CheckAnswer(req CheckAnswerReq) (bool, error) {
	return q.r.CheckAnswer(req)
}

func (q *quesSvc) GetHint(questionID string) (string, error) {
	return q.r.GetHint(questionID)
}
