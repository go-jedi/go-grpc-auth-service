package user

import "context"

func (s *serv) Exists(ctx context.Context, username string, email string) (bool, error) {
	return s.userRepository.Exists(ctx, username, email)
}
