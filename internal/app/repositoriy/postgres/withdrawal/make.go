package withdrawalrepositorypostgres

import (
	"context"
)

func (w *withdrawalPostgresRepository) Make(ctx context.Context, orderNum string, userID int, amount float64) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query := `UPDATE user_balance 
			  SET current=user_balance.current-$1,
    		  withdrawn=user_balance.withdrawn+$1
			  WHERE user_id=$2`

	_, err = tx.ExecContext(ctx, query, amount, userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO withdrawals (order_number, user_id, amount) VALUES ($1, $2, $3)`, orderNum, userID, amount)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
