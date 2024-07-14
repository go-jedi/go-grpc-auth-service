package user

import "context"

func (s *serv) ExistsUsername(ctx context.Context, username string) (bool, error) {
	return s.userRepository.ExistsUsername(ctx, username)
}
