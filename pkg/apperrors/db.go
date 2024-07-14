package apperrors

import "errors"

var (
	ErrNoRowsWereUpdate = errors.New("no rows were updated")
)
