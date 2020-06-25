package questions

import (
	"fmt"
	"strconv"

	"github.com/Devil39/enigma/pkg"
	"github.com/Devil39/enigma/pkg/entities"
	"github.com/Devil39/enigma/pkg/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type repo struct {
	db *sqlx.DB
}

//MakeNewQuestionRepo takes and instance of db and returns a new instance of question repository
func MakeNewQuestionRepo(db *sqlx.DB) Repository {
	return &repo{db: db}
}

//AddQuestion adds a question to the database
func (r *repo) AddQuestion(req AddQuestionReq) error {
	hashedAnswer, err := getHashedAnswer(req.Answer)
	if err != nil {
		return err
	}

	q := fmt.Sprintf("insert into questions(title, description, currpoints, solvecount, answer) values('%v', '%v', 0, 0, '%v');", req.Title, req.Desc, hashedAnswer)

	_, err = r.db.Exec(q)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAllQuestions() ([]entities.Question, error) {
	var questions []entities.Question
	q := fmt.Sprint("select * from questions;")

	rows, err := r.db.Queryx(q)
	for rows.Next() {
		var ques models.QuestionModel
		err := rows.StructScan(&ques)
		if err != nil {
			return nil, err
		}
		questions = append(questions, ques.ToEntity())
	}
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *repo) CheckAnswer(req CheckAnswerReq) (bool, error) {
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return false, pkg.ErrWrongFormat
	}

	q := fmt.Sprintf("select * from questions where questionid = '%d'", id)

	var questionModel models.QuestionModel
	rows, err := r.db.Queryx(q)
	if err != nil || rows.Err() != nil {
		return false, err
	}

	for rows.Next() {
		err = rows.StructScan(&questionModel)
		if err != nil {
			return false, err
		}
	}

	if !compareAnswers(questionModel.Answer, req.Answer) {
		return false, nil
	}

	return true, nil
}

//GetHint takes question id and returns hint of that question
func (r *repo) GetHint(questionID string) (string, error) {
	id, err := strconv.Atoi(questionID)
	if err != nil {
		return "", pkg.ErrWrongFormat
	}

	q := fmt.Sprintf("select * from questions where questionid = '%d'", id)

	var questionModel models.QuestionModel
	rows, err := r.db.Queryx(q)
	if err != nil || rows.Err() != nil {
		return "", err
	}

	cnt := 0
	for rows.Next() {
		err = rows.StructScan(&questionModel)
		if err != nil {
			return "", err
		}
		cnt++
	}

	if cnt != 1 {
		return "", pkg.ErrInvalidQuestionID
	}

	return "hint", nil
}

func getHashedAnswer(answer string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(answer), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func compareAnswers(hashedAnswer string, answer string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedAnswer), []byte(answer))
	return err == nil
}
