package car

import "context"

var (
	deleteQuery = `DELETE FROM car WHERE reg_num = $1`
)

func (r *Repository) DeleteCar(ctx context.Context, regNum string) error {
	_, err := r.db.Exec(ctx, deleteQuery, regNum)
	return err
}
