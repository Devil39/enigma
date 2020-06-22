package user

//Repository represents the user repository
type Repository interface {
	CreateUser() error
	Login() error
}
