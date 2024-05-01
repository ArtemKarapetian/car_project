package car

import (
	"context"
	"github.com/jackc/pgconn"
)

type Repository struct {
	db QueryEngine
}

type QueryEngine interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

func NewRepository(db QueryEngine) *Repository {
	return &Repository{db: db}
}
