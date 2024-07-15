package redis

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-jedi/auth/config"
	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/pkg/postgres"
	"github.com/redis/go-redis/v9"
)

const (
	PrefixUser = "user:"
)

type Redis struct {
	User *User

	// DurCacheUpdate duration in minutes to update cache
	durCacheUpdate int
	db             *postgres.Postgres
}

func NewRedis(ctx context.Context, cfg config.RedisConfig, db *postgres.Postgres) (*Redis, error) {
	r := &Redis{
		durCacheUpdate: cfg.DurCacheUpdate,
		db:             db,
	}

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

	// handle cache add/update cache in goroutine
	go r.handleCache(ctx)

	return r, nil
}

func (r *Redis) handleCache(ctx context.Context) {
	if err := r.cacheUserData(ctx); err != nil {
		log.Println(err)
	}

	for {
		time.Sleep(time.Duration(r.durCacheUpdate) * time.Minute)
		if err := r.reCachedUserData(ctx); err != nil {
			log.Println(err)
		}
	}
}

// cacheUserData cache users data.
func (r *Redis) cacheUserData(ctx context.Context) error {
	usrs, err := r.loadUsers(ctx)
	if err != nil {
		return err
	}

	for _, v := range usrs {
		key := strconv.FormatInt(v.ID, 10)

		if err := r.User.Set(ctx, key, v, 0); err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}

// reCachedUserData reload users data.
func (r *Redis) reCachedUserData(ctx context.Context) error {
	// get all users in db
	udb, err := r.loadUsers(ctx)
	if err != nil {
		return err
	}

	// set users from db to cache
	for _, u := range udb {
		key := strconv.FormatInt(u.ID, 10)

		if err := r.User.Set(ctx, key, u, 0); err != nil {
			log.Println(err)
			continue
		}
	}

	// get all users in cache
	uc, err := r.User.All(ctx)
	if err != nil {
		return err
	}

	// go through the cache and check the presence of each key
	// in the cache for presence from the database
	for _, u := range uc {
		ie, err := r.existsUserInDB(ctx, u.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		if !ie {
			key := strconv.FormatInt(u.ID, 10)
			if err := r.User.Del(ctx, key); err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

// loadUsers load users from db.
func (r *Redis) loadUsers(ctx context.Context) ([]user.User, error) {
	q := `SELECT * FROM users;`

	rows, err := r.db.Pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usrs []user.User
	for rows.Next() {
		var u user.User

		err := rows.Scan(
			&u.ID, &u.Username, &u.FullName, &u.Email,
			&u.PasswordHash, &u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		usrs = append(usrs, u)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return usrs, nil
}

// existsUserInDB check exists user in db.
func (r *Redis) existsUserInDB(ctx context.Context, id int64) (bool, error) {
	ie := false

	q := `SELECT EXISTS(SELECT *  FROM users WHERE id = $1);`

	if err := r.db.Pool.QueryRow(ctx, q, id).Scan(&ie); err != nil {
		return ie, err
	}

	return ie, nil
}
