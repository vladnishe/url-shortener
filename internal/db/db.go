package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
)

func NewPOSTGRES(dsn string) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
