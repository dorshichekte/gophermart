package orderrepositorypostgres

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/order"
	model "gophermarket/internal/app/repositoriy/model/order"
)

func (o *orderPostgresRepository) GetNew(ctx context.Context) ([]model.Order, error) {
	query := `SELECT * 
			  FROM orders
			  WHERE status 
			  IN ($1, $2, $3) LIMIT $4;`

	rows, err := o.db.QueryContext(ctx, query, entity.StatusNew, entity.StatusRegistered, entity.StatusProcessing, 1000)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err = rows.Scan(&order.ID, &order.Number, &order.Status, &order.UserID, &order.Accrual, &order.Active, &order.UploadAt, &order.ModifiedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
