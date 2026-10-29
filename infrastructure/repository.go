package infrastructure

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type QueryOptions struct {
	offset int
	limit  int
}

func NewQueryOptions(offset, limit int) QueryOptions {
	return QueryOptions{
		offset, limit,
	}
}

type DatabaseConnector interface {
	Ping(ctx context.Context) error
	RunQuery(ctx context.Context, result any, query string, args ...any) error
	RunCommand(ctx context.Context, operationFn func(*sqlx.Tx) error) error
}
