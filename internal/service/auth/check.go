package auth

import (
	"context"
	"log"
	"strconv"

	"github.com/go-jedi/auth/internal/domain/auth"
	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) Check(ctx context.Context, dto auth.CheckDTO) (auth.CheckResp, error) {
	u, err := s.checkGetFromCacheOrRepo(ctx, dto.ID)
	if err != nil {
		return auth.CheckResp{}, err
	}

	// check verify token
	vr, err := s.jwt.Verify(u.ID, u.Username, dto.Token)
	if err != nil {
		return auth.CheckResp{}, err
	}

	return auth.CheckResp{
		ID:       vr.ID,
		Username: vr.Username,
		Token:    dto.Token,
		ExpAt:    vr.ExpAt,
	}, nil
}

// checkGetFromCacheOrRepo get user from cache.
func (s *serv) checkGetFromCacheOrRepo(ctx context.Context, id int64) (user.User, error) {
	key := strconv.FormatInt(id, 10)

	u, err := s.cache.User.Get(ctx, key)
	if err != nil {
		log.Println(err)
		return s.userRepository.GetByID(ctx, id)
	}

	return u, nil
}
