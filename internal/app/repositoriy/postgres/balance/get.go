package balancerepositorypostgres

import (
	"context"

	model "gophermarket/internal/app/repositoriy/model/balance"
)

func (b *balancePostgresRepository) Get(ctx context.Context, userID int) (model.Balance, error) {
	var balance model.Balance

	query := `
		SELECT * FROM balance
		WHERE user_id = $1;
	`
	err := b.db.QueryRowContext(ctx, query, userID).Scan(&balance.ID, &balance.UserID, &balance.Current, &balance.Withdrawn, &balance.CreatedAt)
	if err != nil {
		return balance, err
	}

	return balance, nil
}
