package orderrepositorypostgres

import (
	"context"

	model "gophermarket/internal/app/repositoriy/model/order"
)

func (o *orderPostgresRepository) GetAll(ctx context.Context, userID int) ([]model.Order, error) {
	query := `SELECT *
			  FROM orders
			  WHERE user_id = $1
			  ORDER BY upload_at DESC
	`
	rows, err := o.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		if err = rows.Scan(&o.ID, &o.Number, &o.Status, &o.UserID, &o.Accrual, &o.Active, &o.UploadAt, &o.ModifiedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
