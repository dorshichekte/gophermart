package userrepositorypostgres

import (
	"context"

	model "gophermarket/internal/app/repositoriy/model/user"
)

func (r *userPostgresRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	query := `
			SELECT id, login, password 
			FROM users 
			WHERE login = $1;
			`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
