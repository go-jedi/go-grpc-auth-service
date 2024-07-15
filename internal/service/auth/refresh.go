package auth

import (
	"context"
	"log"
	"strconv"

	"github.com/go-jedi/auth/internal/domain/auth"
	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) Refresh(ctx context.Context, dto auth.RefreshDTO) (auth.RefreshResp, error) {
	u, err := s.refreshGetFromCacheOrRepo(ctx, dto.ID)
	if err != nil {
		return auth.RefreshResp{}, err
	}

	// check verify token
	vr, err := s.jwt.Verify(u.ID, u.Username, dto.RefreshToken)
	if err != nil {
		return auth.RefreshResp{}, err
	}

	// generate access, refresh tokens
	gr, err := s.jwt.Generate(vr.ID, vr.Username)
	if err != nil {
		return auth.RefreshResp{}, err
	}

	return auth.RefreshResp{
		AccessToken:  gr.AccessToken,
		RefreshToken: gr.RefreshToken,
		AccessExpAt:  gr.AccessExpAt,
		RefreshExpAt: gr.RefreshExpAt,
	}, nil
}

// refreshGetFromCacheOrRepo get user from cache.
func (s *serv) refreshGetFromCacheOrRepo(ctx context.Context, id int64) (user.User, error) {
	key := strconv.FormatInt(id, 10)

	u, err := s.cache.User.Get(ctx, key)
	if err != nil {
		log.Println(err)
		return s.userRepository.GetByID(ctx, id)
	}

	return u, nil
}
