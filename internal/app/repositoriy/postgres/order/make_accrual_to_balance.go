package orderrepositorypostgres

import (
	"context"
	"time"

	model "gophermarket/internal/app/repositoriy/model/order"
)

func (o *orderPostgresRepository) MakeAccrualToBalance(ctx context.Context, order model.Order) error {
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	queryUpdateOrders := `UPDATE orders
			  SET status = $1, accrual = $2, modified_at = $3 
			  WHERE number = $4`
	_, err = tx.ExecContext(ctx, queryUpdateOrders, order.Status, order.Accrual, time.Now(), order.Number)
	if err != nil {
		return err
	}

	queryUpdateBalance := `UPDATE user_balance 
 						   SET curren=user_balance.current+$1
 						   WHERE user_id=$2`
	_, err = tx.ExecContext(ctx, queryUpdateBalance, order.Accrual, order.UserID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
