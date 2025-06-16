package user_usecase

import (
	"go.uber.org/zap"

	user_repository "gophermarket/internal/app/domain/repository/user"
	"gophermarket/internal/libs/auth"
	"gophermarket/internal/libs/hasher"
)

type UserUseCase struct {
	hasher         hasher.Hasher
	logger         *zap.Logger
	auth           auth.Auth
	userRepository user_repository.UserRepository
}
