package postgres

import (
	"car_project/internal/config"
	"car_project/internal/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

func NewDb(ctx context.Context, envPath string) (*Database, error) {
	if err := config.Load(envPath); err != nil {
		return nil, errors.Wrap(err, model.ErrorEnvLoad)
	}
	dsn, err := config.GeneratePostgresDSN()
	if err != nil {
		return nil, errors.Wrap(err, model.ErrorDSNGenerate)

	}
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return newDatabase(pool), nil
}
