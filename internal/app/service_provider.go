package app

import (
	"car_project/internal/db/postgres"
	"car_project/internal/repository"
)

type serviceProvider struct {
	db   *postgres.Database
	repo repository.CarRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}
