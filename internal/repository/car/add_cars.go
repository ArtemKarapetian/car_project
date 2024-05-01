package car

import "context"

var (
	addCarsQuery = `INSERT INTO car (reg_num) VALUES ($1)`
)

func (r *Repository) AddCars(ctx context.Context, regNums []string) error {
	_, err := r.db.Exec(ctx, addCarsQuery, regNums)
	return err
}
