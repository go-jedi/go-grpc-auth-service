package redis

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/pkg/apperrors"
	"github.com/redis/go-redis/v9"
	"github.com/vmihailenco/msgpack/v5"
)

type User struct {
	client *redis.Client
}

func NewUser(client *redis.Client) *User {
	return &User{client: client}
}

// Set create/update value by key.
func (u *User) Set(ctx context.Context, key string, val user.User, expiration time.Duration) error {
	b, err := msgpack.Marshal(val)
	if err != nil {
		return err
	}

	p := u.client.Pipeline()
	p.Set(ctx, PrefixUser+key, b, expiration)
	if _, err = p.Exec(ctx); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// All get all users.
func (u *User) All(ctx context.Context) ([]user.User, error) {
	keys, err := u.client.Keys(ctx, PrefixUser+"*").Result()
	if err != nil {
		return nil, err
	}

	usrs := make([]user.User, 0, len(keys))

	for _, k := range keys {
		b, err := u.client.Get(ctx, k).Bytes()
		if err != nil {
			log.Println(err)
			continue
		}

		var usr user.User
		if err := msgpack.Unmarshal(b, &usr); err != nil {
			log.Println(err)
			continue
		}

		usrs = append(usrs, usr)
	}

	return usrs, nil
}

// Get value by key.
func (u *User) Get(ctx context.Context, key string) (user.User, error) {
	b, err := u.client.Get(ctx, PrefixUser+key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return user.User{}, apperrors.ErrCacheKeyNotExists
		}
		log.Println(err)
		return user.User{}, err
	}

	var usr user.User
	if err := msgpack.Unmarshal(b, &usr); err != nil {
		log.Println(err)
		return user.User{}, err
	}

	return usr, nil
}

// Del key/keys.
func (u *User) Del(ctx context.Context, keys ...string) error {
	p := u.client.Pipeline()

	for _, key := range keys {
		p.Del(ctx, PrefixUser+key)
	}

	_, err := p.Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
