package models

import (
	"strings"

	"github.com/Devil39/enigma/pkg/entities"
)

//UserModel represents user model struct
type UserModel struct {
	UUID            string `db:"uid"`
	EmailID         string `db:"emailid"`
	Password        string `db:"password"`
	SolvedQuestions string `db:"questionssolved"`
	HintsUsed       string `db:"hintsused"`
}

//ToEntity converts given user model to user entity
func (userModel *UserModel) ToEntity() entities.User {
	return entities.User{
		UUID:            userModel.UUID,
		EmailID:         userModel.EmailID,
		SolvedQuestions: strings.Split(userModel.SolvedQuestions, ","),
		HintsUsed:       strings.Split(userModel.HintsUsed, ","),
	}
}
