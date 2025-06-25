package orderrepositorypostgres

import (
	"context"

	model "gophermarket/internal/app/repositoriy/model/order"
)

func (o *orderPostgresRepository) GetByNumber(ctx context.Context, orderNumber string) (model.Order, error) {
	query := `SELECT *
			  FROM orders
			  WHERE number = $1
	`
	row := o.db.QueryRowContext(ctx, query, orderNumber)

	var order model.Order
	err := row.Scan(&order.ID, &order.Number, &order.Status, &order.UserID, &order.Accrual, &order.Active, &order.UploadAt, &order.ModifiedAt)
	if err != nil {
		return order, err
	}

	return order, nil
}
