package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (r *repo) Create(ctx context.Context, dto user.CreateDTO) (user.User, error) {
	var nu user.User

	q := `
		INSERT INTO users(
			username,
		    full_name,
		    email,
		    password_hash
		) VALUES($1, $2, $3, $4)
		RETURNING id, username, full_name, email, password_hash, created_at, updated_at;
	`

	if err := r.db.Pool.QueryRow(ctx, q, dto.Username, dto.FullName, dto.Email, dto.Password).Scan(
		&nu.ID, &nu.Username, &nu.FullName, &nu.Email,
		&nu.PasswordHash, &nu.CreatedAt, &nu.UpdatedAt,
	); err != nil {
		return user.User{}, err
	}

	return nu, nil
}
