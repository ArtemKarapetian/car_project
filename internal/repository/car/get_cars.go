package car

import (
	"car_project/internal/model"
	"car_project/internal/repository/car/converter"
	repomodel "car_project/internal/repository/car/model"
	"context"
)

var (
	getCarsQuery = `
		SELECT c.reg_num, c.mark, c.model, c.year, p.name, p.surname, p.patronymic
		FROM cars c
		JOIN people p ON c.owner_id = p.id
		WHERE ($1::text IS NULL OR c.reg_num = $1)
		AND ($2::text IS NULL OR c.mark = $2)
		AND ($3::text IS NULL OR c.model = $3)
		AND ($4::int IS NULL OR c.year = $4 OR $4 = 0)
		AND ($5::text IS NULL OR p.name = $5)
		AND ($6::text IS NULL OR p.surname = $6)
		AND ($7::text IS NULL OR p.patronymic = $7)
		LIMIT $8 OFFSET $9
	`
)

func (r *repository) GetCars(ctx context.Context, filter *model.Car, limit, offset int) ([]*model.Car, error) {
	repoFilter := converter.ToRepoFromCar(filter)
	cars := make([]*repomodel.Car, 0)
	err := r.db.Select(ctx, cars, getCarsQuery, repoFilter.RegNum, repoFilter.Mark, repoFilter.Model, repoFilter.Year, repoFilter.Owner.Name, repoFilter.Owner.Surname, repoFilter.Owner.Patronymic, limit, offset)
	if err != nil {
		return nil, err
	}

	resultCars := make([]*model.Car, 0, len(cars))
	for _, car := range cars {
		resultCars = append(resultCars, converter.ToCarFromRepo(car))
	}

	return resultCars, nil
}
