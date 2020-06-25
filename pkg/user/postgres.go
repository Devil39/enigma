package user

import (
	"fmt"

	"github.com/Devil39/enigma/pkg"
	"github.com/Devil39/enigma/pkg/entities"
	"github.com/Devil39/enigma/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

type repo struct {
	db *sqlx.DB
}

//AuthRequest depicts the login request structure
type AuthRequest struct {
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}

//NewPostgresRepo returns a new user repository
func NewPostgresRepo(db *sqlx.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateUser(req AuthRequest) (entities.User, error) {
	userUUID := uuid.NewV4()

	hashedPassword, err := getHashedPassword(req.Password)
	if err != nil {
		return entities.User{}, err
	}

	q := fmt.Sprintf("insert into users values('%v', '%v', '%v', '', '');", userUUID, req.EmailID, hashedPassword)

	_, err = r.db.Exec(q)
	if err != nil {
		return entities.User{}, err
	}

	user := models.UserModel{
		UUID:            userUUID.String(),
		EmailID:         req.EmailID,
		Password:        hashedPassword,
		SolvedQuestions: "",
		HintsUsed:       "",
	}

	return user.ToEntity(), nil
}

func (r *repo) Login(req AuthRequest) (entities.User, error) {

	q := fmt.Sprintf("select * from users where emailid = '%v'", req.EmailID)

	var userModel models.UserModel
	rows, err := r.db.Queryx(q)
	if err != nil || rows.Err() != nil {
		return entities.User{}, pkg.ErrInvalidEmailAndPass
	}

	for rows.Next() {
		err = rows.StructScan(&userModel)
		if err != nil {
			return entities.User{}, err
		}
	}

	if !comparePassword(userModel.Password, req.Password) {
		return entities.User{}, pkg.ErrInvalidEmailAndPass
	}

	return userModel.ToEntity(), nil
}

func (r *repo) AddSolvedQuestion(emailID, questionID string) error {

	q := fmt.Sprintf("select * from users where emailid = '%v'", emailID)

	var userModel models.UserModel
	rows, err := r.db.Queryx(q)
	if err != nil || rows.Err() != nil {
		return pkg.ErrInvalidEmailAndPass
	}

	for rows.Next() {
		err = rows.StructScan(&userModel)
		if err != nil {
			return err
		}

		q := fmt.Sprintf("update users set questionssolved = '%v' where emailid = '%v'", userModel.SolvedQuestions+questionID+",", userModel.EmailID)

		_, err := r.db.Exec(q)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repo) AddHintUsed(emailID, questionID string) error {
	q := fmt.Sprintf("select * from users where emailid = '%v'", emailID)

	var userModel models.UserModel
	rows, err := r.db.Queryx(q)
	if err != nil || rows.Err() != nil {
		return pkg.ErrInvalidEmailAndPass
	}

	cnt := 0
	for rows.Next() {
		err = rows.StructScan(&userModel)
		if err != nil {
			return err
		}

		q := fmt.Sprintf("update users set hintsused = '%v' where emailid = '%v'", userModel.SolvedQuestions+questionID+",", userModel.EmailID)

		_, err := r.db.Exec(q)
		if err != nil {
			return err
		}

		cnt++
	}

	if cnt != 1 {
		return pkg.ErrInvalidUserID
	}

	return nil
}

func getHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
