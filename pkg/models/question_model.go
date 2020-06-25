package models

import "github.com/Devil39/enigma/pkg/entities"

//QuestionModel acts as a structure for response from database
type QuestionModel struct {
	ID         string  `db:"questionid"`
	Title      string  `db:"title"`
	Desc       string  `db:"description"`
	CurrPoints float64 `db:"currpoints"`
	SolveCount float64 `db:"solvecount"`
	Answer     string  `db:"answer"`
}

//ToEntity converts question model to an entity
func (q *QuestionModel) ToEntity() entities.Question {
	return entities.Question{
		ID:    q.ID,
		Title: q.Title,
		Desc:  q.Desc,
	}
}
