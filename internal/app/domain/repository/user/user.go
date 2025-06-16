package userrepository

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/user"
	model "gophermarket/internal/app/repositoriy/model/user"
)

type UserRepository interface {
	Register(ctx context.Context, user entity.User) (int, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
}
