package redis

import (
	"context"
	"time"

	"github.com/go-jedi/auth/config"
	"github.com/redis/go-redis/v9"
)

const (
	PrefixUser = "user:"
)

type Redis struct {
	User *User
}

func NewRedis(ctx context.Context, cfg config.RedisConfig) (*Redis, error) {
	r := &Redis{}

	c := redis.NewClient(&redis.Options{
		Addr:            cfg.Addr,
		Password:        cfg.Password,
		DB:              cfg.DB,
		DialTimeout:     time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:     time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(cfg.WriteTimeout) * time.Second,
		PoolSize:        10,
		MinIdleConns:    3,
		PoolTimeout:     time.Duration(cfg.PoolTimeout) * time.Second,
		TLSConfig:       nil,
		MaxRetries:      3,
		MinRetryBackoff: time.Duration(cfg.MinRetryBackoff) * time.Millisecond,
		MaxRetryBackoff: time.Duration(cfg.MaxRetryBackoff) * time.Millisecond,
	})

	_, err := c.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	r.User = NewUser(c)

	return r, nil
}
