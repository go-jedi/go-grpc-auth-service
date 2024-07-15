package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (r *repo) All(ctx context.Context) ([]user.User, error) {
	q := `
		SELECT * 
		FROM users;
	`

	rows, err := r.db.Pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usrs []user.User
	for rows.Next() {
		var u user.User

		err := rows.Scan(
			&u.ID, &u.Username, &u.FullName, &u.Email,
			&u.PasswordHash, &u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		usrs = append(usrs, u)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return usrs, nil
}
