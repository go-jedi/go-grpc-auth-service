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
