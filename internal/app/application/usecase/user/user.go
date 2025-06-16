package user_usecase

import (
	"go.uber.org/zap"

	user_repository "gophermarket/internal/app/domain/repository/user"
	"gophermarket/internal/libs/auth"
	"gophermarket/internal/libs/hasher"
)

func New(logger *zap.Logger, hasher hasher.Hasher, auth auth.Auth, userRepository user_repository.UserRepository) *UserUseCase {
	return &UserUseCase{hasher: hasher, logger: logger, auth: auth, userRepository: userRepository}
}
