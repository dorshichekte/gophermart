package order_repository_postgres

import (
	"context"
	customerror "gophermarket/internal/error"

	entity "gophermarket/internal/app/domain/entity/order"
)

func (o *orderPostgresRepository) Upload(ctx context.Context, order entity.Order) error {
	query := `INSERT INTO orders(number, user_id, status)
			  VALUES ($1, $2, $3)	
	`
	res, err := o.db.ExecContext(ctx, query, order.Number, order.UserID, order.Status)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return customerror.New("upload order: no rows affected")
	}

	return nil
}
