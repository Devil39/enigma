package user

import (
	"database/sql"
	"fmt"
	//"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sql.DB
}

//NewPostgresRepo returns a new user repository
func NewPostgresRepo(db *sql.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateUser() error {
	fmt.Println("Creating User")
	//r.db.Exec(fmt.Sprintf("INSERT INTO TABLE users VALUES()"))
	return nil
}

func (r *repo) Login() error {
	fmt.Println("Login")
	return nil
}
