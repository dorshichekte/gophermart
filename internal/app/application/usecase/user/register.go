package user_usecase

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/user"
)

func (uCase *UserUseCase) Register(ctx context.Context, login, password string) (int, error) {
	user, _ := uCase.userRepository.GetByLogin(ctx, login)
	if user != nil {
		return 0, ErrLoginAlreadyTaken
	}

	hashedPassword, err := uCase.hasher.Hash(password)
	if err != nil {
		return 0, err
	}

	userEntity := entity.NewUser(login, hashedPassword)

	userID, err := uCase.userRepository.Register(ctx, userEntity)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
