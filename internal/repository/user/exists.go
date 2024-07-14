package user

import (
	"context"
)

func (r *repo) Exists(ctx context.Context, username string, email string) (bool, error) {
	ie := false

	q := `
		SELECT EXISTS(
			SELECT * 
			FROM users 
			WHERE username = $1
			OR email = $2
		)
	`

	if err := r.db.Pool.QueryRow(ctx, q, username, email).Scan(&ie); err != nil {
		return ie, err
	}

	return ie, nil
}
