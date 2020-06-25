package user

import "github.com/Devil39/enigma/pkg/entities"

//Repository represents the user repository
type Repository interface {
	CreateUser(AuthRequest) (entities.User, error)
	Login(AuthRequest) (entities.User, error)
	AddSolvedQuestion(string, string) error
	AddHintUsed(string, string) error
}
