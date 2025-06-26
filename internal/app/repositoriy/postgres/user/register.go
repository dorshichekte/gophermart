package userrepositorypostgres

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/user"
	model "gophermarket/internal/app/repositoriy/model/user"
)

func (r *userPostgresRepository) Register(ctx context.Context, user entity.User) (int, error) {
	dbUser := model.NewUser(user)

	trs, beginErr := r.db.BeginTx(ctx, nil)
	if beginErr != nil {
		return 0, beginErr
	}
	defer func() {
		_ = trs.Rollback()
	}()

	userQuery := `
			INSERT INTO users (login, password)
			VALUES ($1, $2)
			RETURNING id;
			`
	var id int
	err := trs.QueryRowContext(ctx, userQuery, dbUser.Login, dbUser.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	userBalanceQuery := `
		INSERT INTO user_balance (user_id)
		VALUES ($1);
	`
	_, insertErr := trs.ExecContext(ctx, userBalanceQuery, id)
	if insertErr != nil {
		return 0, insertErr
	}

	if transactionErr := trs.Commit(); transactionErr != nil {
		return 0, transactionErr
	}

	return id, nil
}
