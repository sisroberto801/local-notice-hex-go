package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connector interface {
	Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error)
}

type SqlConnector struct{}

func NewSqlConnector() Connector {
	return &SqlConnector{}
}

func (p *SqlConnector) Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, databaseURL)
}
