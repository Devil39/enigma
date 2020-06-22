package models

//UserModel represents user model struct
type UserModel struct {
	uuid            string
	emailID         string
	solvedQuestions []string
	hintsUsed       []string
}
