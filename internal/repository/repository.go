package repository

import (
	"car_project/internal/model"
	"context"
)

type CarRepository interface {
	// GetCars получает данных с фильтрацией по всем полям и пагинацией
	GetCars(ctx context.Context, filter *model.Car, limit, offset int) ([]*model.Car, error)

	// DeleteCar удаляет по идентификатору
	DeleteCar(ctx context.Context, regNum string) error

	// UpdateCar измененяет одного или нескольких полей по идентификатору
	UpdateCar(ctx context.Context, regNum string, car *model.Car) error

	// AddCars добавляет новых автомобилей
	AddCars(ctx context.Context, regNums []string) error
}
