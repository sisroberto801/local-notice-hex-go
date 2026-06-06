package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connector interface {
	Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error)
}

type PostgreSQLConnector struct{}

func NewPostgreSQLConnector() Connector {
	return &PostgreSQLConnector{}
}

func (p *PostgreSQLConnector) Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, databaseURL)
}
