package user

import "github.com/Devil39/enigma/pkg/entities"

//Service represents user service
type Service interface {
	CreateUser(AuthRequest) (entities.User, error)
	Login(AuthRequest) (entities.User, error)
	AddSolvedQuestion(string, string) error
	AddHintUsed(string, string) error
}

type userSvc struct {
	repo Repository
}

//NewUserService returns a new user service
func NewUserService(r Repository) Service {
	return &userSvc{repo: r}
}

func (uS *userSvc) CreateUser(req AuthRequest) (entities.User, error) {
	return uS.repo.CreateUser(req)
}

func (uS *userSvc) Login(req AuthRequest) (entities.User, error) {
	return uS.repo.Login(req)
}

func (uS *userSvc) AddSolvedQuestion(emailID string, questionID string) error {
	return uS.repo.AddSolvedQuestion(emailID, questionID)
}

func (uS *userSvc) AddHintUsed(emailID string, questionID string) error {
	return uS.repo.AddHintUsed(emailID, questionID)
}
