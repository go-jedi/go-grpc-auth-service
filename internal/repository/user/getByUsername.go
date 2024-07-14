package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (r *repo) GetByUsername(ctx context.Context, username string) (user.User, error) {
	var u user.User

	q := `
		SELECT *
		FROM users
		WHERE username = $1
	`

	if err := r.db.Pool.QueryRow(ctx, q, username).Scan(
		&u.ID, &u.Username, &u.FullName, &u.Email,
		&u.PasswordHash, &u.Deleted, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return user.User{}, err
	}

	return u, nil
}
