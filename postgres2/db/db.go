package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context, connStr string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	return conn, err
}
