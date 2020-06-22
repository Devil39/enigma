package user

import "fmt"

type repo struct {
	n int
}

//NewPostgresRepo returns a new user repository
func NewPostgresRepo(n int) Repository {
	return &repo{n: n}
}

func (r *repo) CreateUser() error {
	fmt.Println("Creating User")
	return nil
}

func (r *repo) Login() error {
	fmt.Println("Login")
	return nil
}
