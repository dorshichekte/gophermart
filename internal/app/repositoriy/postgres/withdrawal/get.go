package withdrawalrepositorypostgres

import (
	"context"
	model "gophermarket/internal/app/repositoriy/model/withdrawal"
)

func (w *withdrawalPostgresRepository) Get(ctx context.Context, userID int) ([]model.Withdrawal, error) {
	query := `SELECT * 
			  FROM withdrawals
			  WHERE user_id = $1
			  ORDER BY created_at DESC`

	rows, err := w.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var withdrawals []model.Withdrawal
	for rows.Next() {
		var wd model.Withdrawal
		if rowErr := rows.Scan(&wd.ID, &wd.OrderNumber, &wd.UserID, &wd.Amount, &wd.CreatedAt); rowErr != nil {
			return nil, rowErr
		}
		withdrawals = append(withdrawals, wd)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return withdrawals, nil
}
