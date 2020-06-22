package user

//Service represents user service
type Service interface {
	CreateUser() error
	Login() error
}

type userSvc struct {
	repo Repository
}

//NewUserService returns a new user service
func NewUserService(r Repository) Service {
	return &userSvc{repo: r}
}

func (uS *userSvc) CreateUser() error {
	return uS.repo.CreateUser()
}

func (uS *userSvc) Login() error {
	return uS.repo.Login()
}
