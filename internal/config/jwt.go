package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	refreshTokenSecretKeyEnvName  = "REFRESH_TOKEN_SECRET_KEY"
	accessTokenSecretKeyEnvName   = "ACCESS_TOKEN_SECRET_KEY"
	refreshTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpirationEnvName  = "ACCESS_TOKEN_EXPIRATION"
)

type JWTConfig interface {
	RefreshTokenSecretKey() string
	AccessTokenSecretKey() string
	RefreshTokenExpiration() string
	AccessTokenExpiration() string
}

type jwtConfig struct {
	refreshTokenSecretKey  string
	accessTokenSecretKey   string
	refreshTokenExpiration string
	accessTokenExpiration  string
}

func NewJWTConfig() (JWTConfig, error) {
	refreshTokenSecretKey := os.Getenv(refreshTokenSecretKeyEnvName)
	if len(refreshTokenSecretKey) == 0 {
		return nil, errors.New("refresh token secret key not found")
	}

	accessTokenSecretKey := os.Getenv(accessTokenSecretKeyEnvName)
	if len(accessTokenSecretKey) == 0 {
		return nil, errors.New("access token secret key not found")
	}

	refreshTokenExpiration := os.Getenv(refreshTokenExpirationEnvName)
	if len(refreshTokenExpiration) == 0 {
		return nil, errors.New("refresh token expiration key not found")
	}

	accessTokenExpiration := os.Getenv(accessTokenExpirationEnvName)
	if len(accessTokenExpiration) == 0 {
		return nil, errors.New("access token expiration key not found")
	}

	return &jwtConfig{
		refreshTokenSecretKey:  refreshTokenSecretKey,
		accessTokenSecretKey:   accessTokenSecretKey,
		refreshTokenExpiration: refreshTokenExpiration,
		accessTokenExpiration:  accessTokenExpiration,
	}, nil
}

func (cfg *jwtConfig) RefreshTokenSecretKey() string {
	return cfg.refreshTokenSecretKey
}

func (cfg *jwtConfig) AccessTokenSecretKey() string {
	return cfg.accessTokenSecretKey
}

func (cfg *jwtConfig) RefreshTokenExpiration() string {
	return cfg.refreshTokenExpiration
}

func (cfg *jwtConfig) AccessTokenExpiration() string {
	return cfg.accessTokenExpiration
}
