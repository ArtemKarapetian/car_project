package car

import (
	"car_project/internal/model"
	"car_project/internal/repository/car/converter"
	"context"
)

var (
	updateQuery = "UPDATE car SET mark = $1, model = $2, year = $3 WHERE reg_num = $5"
)

func (r *repository) UpdateCar(ctx context.Context, regNum string, car *model.Car) error {
	repoCar := converter.ToRepoFromCar(car)
	_, err := r.db.Exec(ctx, updateQuery, repoCar.Mark, repoCar.Model, repoCar.Year, regNum)
	return err
}
