package userusecase

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (uCase *UserUseCase) Login(ctx context.Context, login, password string) (int, error) {
	user, err := uCase.userRepository.GetByLogin(ctx, login)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return 0, ErrUserNotFound
		}
		return 0, err
	}

	isCompare := uCase.hasher.Compare(user.Password, password)
	if !isCompare {
		return 0, ErrUserNotFound
	}

	return user.ID, nil
}
