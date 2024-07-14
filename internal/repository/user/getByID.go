package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (r *repo) GetByID(ctx context.Context, id int64) (user.User, error) {
	var u user.User

	q := `
		SELECT *
		FROM users
		WHERE id = $1
	`

	if err := r.db.Pool.QueryRow(ctx, q, id).Scan(
		&u.ID, &u.Username, &u.FullName, &u.Email,
		&u.PasswordHash, &u.Deleted, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return user.User{}, err
	}

	return u, nil
}
