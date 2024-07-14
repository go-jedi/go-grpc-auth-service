package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (r *repo) Update(ctx context.Context, dto user.UpdateDTO) (user.User, error) {
	var uu user.User

	q := `
		UPDATE users SET
		    username = $1,
		    full_name = $2,
		    email = $3,
		    password_hash = $4
		WHERE id = $5
		RETURNING id, username, full_name, email, password_hash, deleted, created_at, updated_at;
	`

	if err := r.db.Pool.QueryRow(ctx, q, dto.Username, dto.FullName, dto.Email, dto.Password, dto.ID).Scan(
		&uu.ID, &uu.Username, &uu.FullName, &uu.Email,
		&uu.PasswordHash, &uu.Deleted, &uu.CreatedAt, &uu.UpdatedAt,
	); err != nil {
		return user.User{}, err
	}

	return uu, nil
}
