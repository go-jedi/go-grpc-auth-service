package user

import "context"

func (r *repo) ExistsEmail(ctx context.Context, email string) (bool, error) {
	ie := false

	q := `
		SELECT EXISTS(
			SELECT *
			FROM users
			WHERE email = $1
		);
	`

	if err := r.db.Pool.QueryRow(ctx, q, email).Scan(&ie); err != nil {
		return ie, err
	}

	return ie, nil
}
