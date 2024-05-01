package car

import (
	"car_project/internal/model"
	"car_project/internal/repository/car/converter"
	"context"
	"fmt"
	"strings"
)

var (
	updateCarQuery = "UPDATE car SET %s WHERE reg_num = $%d"
)

func (r *repository) UpdateCar(ctx context.Context, regNum string, car *model.Car) error {
	repoCar := converter.ToRepoFromCar(car)

	var updateFields []string
	var updateValues []interface{}

	if repoCar.Mark != "" {
		updateFields = append(updateFields, "mark = $%d")
		updateValues = append(updateValues, repoCar.Mark)
	}
	if repoCar.Model != "" {
		updateFields = append(updateFields, "model = $%d")
		updateValues = append(updateValues, repoCar.Model)
	}
	if repoCar.Year != 0 {
		updateFields = append(updateFields, "year = $%d")
		updateValues = append(updateValues, repoCar.Year)
	}

	if len(updateFields) == 0 {
		return nil
	}

	updateQuery := fmt.Sprintf(updateCarQuery, strings.Join(updateFields, ", "), len(updateValues)+1)
	updateValues = append(updateValues, regNum)

	_, err := r.db.Exec(ctx, updateQuery, updateValues...)
	return err
}
