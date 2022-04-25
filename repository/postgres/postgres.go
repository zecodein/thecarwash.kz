package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/zecodein/thecarwash.kz/configs"
)

func NewPostgresRepository(c *configs.Config) (*pgxpool.Pool, error) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s%s/%s", c.UserDB, c.PasswordDB, c.HostDB, c.PortDB, c.NameDB)

	config, err := pgxpool.ParseConfig(DSN)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	return db, nil
}
