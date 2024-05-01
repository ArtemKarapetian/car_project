package app

import (
	"car_project/internal/db/postgres"
	"car_project/internal/repository"
	carservice "car_project/internal/service/car"
)

type serviceProvider struct {
	db      *postgres.Database
	repo    repository.CarRepository
	handler *carservice.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}
