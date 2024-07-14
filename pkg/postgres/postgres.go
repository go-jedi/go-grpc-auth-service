package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-jedi/auth/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool

	host          string
	user          string
	password      string
	dbName        string
	port          int
	sslMode       string
	poolMaxConns  int
	migrationsDir string
}

func NewPostgres(ctx context.Context, cfg config.PostgresConfig) (*Postgres, error) {
	p := &Postgres{
		host:          cfg.Host,
		user:          cfg.User,
		password:      cfg.Password,
		dbName:        cfg.DBName,
		port:          cfg.Port,
		sslMode:       cfg.SSLMode,
		poolMaxConns:  cfg.PoolMaxConns,
		migrationsDir: cfg.MigrationsDir,
	}

	pool, err := pgxpool.New(ctx, p.generateDsn())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	p.Pool = pool

	if err := p.migrationsUp(); err != nil {
		return nil, err
	}

	return p, nil
}

// generateDsn generate dsn string.
func (p *Postgres) generateDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s pool_max_conns=%d",
		p.host, p.user, p.password, p.dbName, p.port, p.sslMode, p.poolMaxConns,
	)
}

// migrationsUp up migrations for db.
func (p *Postgres) migrationsUp() error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.user, p.password, p.host, p.port, p.dbName, p.sslMode,
	)

	m, err := migrate.New(
		p.migrationsDir,
		dbURL,
	)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}

	return nil
}
