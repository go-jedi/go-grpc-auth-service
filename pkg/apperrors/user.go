package apperrors

import "errors"

var (
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrUsernameAlreadyExists    = errors.New("username already exists")
	ErrEmailAlreadyExists       = errors.New("email already exists")
	ErrUserPasswordNotGenerated = errors.New("password generation error")
	ErrUserIncorrectPassword    = errors.New("incorrect password")
)
