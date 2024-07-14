package apperrors

import "errors"

var (
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrUserPasswordNotGenerated = errors.New("password generation error")
	ErrUserIncorrectPassword    = errors.New("incorrect password")
)
