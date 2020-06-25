package pkg

import "errors"

var (
	//ErrUnauthorized depicts unathorized error
	ErrUnauthorized = errors.New("error: You are unauthorized to access this resource")
	//ErrInvalidEmailAndPass depicts wrong email and password combination
	ErrInvalidEmailAndPass = errors.New("error: Invalid email and password combination")
	//ErrWrongFormat depicts wrong json format sent
	ErrWrongFormat = errors.New("error: Invalid format of data sent")
	//ErrInvalidQuestionID depicts wrong question id is sent
	ErrInvalidQuestionID = errors.New("error: Invalid question ID")
	//ErrInvalidUserID depicts wrong user id is sent
	ErrInvalidUserID = errors.New("error: Invalid user ID")
)
