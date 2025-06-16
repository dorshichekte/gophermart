package withdrawal_repository_postgres

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
		if err := rows.Scan(&wd.ID, &wd.OrderNumber, wd.UserID, &wd.Amount, &wd.CreatedAt); err != nil {
			return nil, err
		}
		withdrawals = append(withdrawals, wd)
	}

	return withdrawals, nil
}
