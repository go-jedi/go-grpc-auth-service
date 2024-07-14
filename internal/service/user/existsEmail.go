package user

import "context"

func (s *serv) ExistsEmail(ctx context.Context, email string) (bool, error) {
	return s.userRepository.ExistsEmail(ctx, email)
}
