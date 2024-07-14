package user

import "context"

func (r *repo) ExistsUsername(ctx context.Context, username string) (bool, error) {
	ie := false

	q := `
		SELECT EXISTS(
			SELECT *
			FROM users
			WHERE username = $1
		);
	`

	if err := r.db.Pool.QueryRow(ctx, q, username).Scan(&ie); err != nil {
		return ie, err
	}

	return ie, nil
}
