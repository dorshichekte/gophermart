package userrepository

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/user"
	model "gophermarket/internal/app/repositoriy/model/user"
)

//go:generate mockgen -package mocks -source user.go -destination ../../mock/user_repository.go UserRepository
type UserRepository interface {
	Register(ctx context.Context, user entity.User) (int, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
}
