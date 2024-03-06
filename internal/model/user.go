package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID                   int64
	Username             string
	Password             string
	CreatedAt            time.Time
	UpdatedAt            sql.NullTime
	PasswordLastChangeAt time.Time
}

type UpdateNameRequest struct {
	ID       int64
	Username string
}

type UpdatePasswordRequest struct {
	ID       int64
	Password string
}
