package model

import (
	entity "gophermarket/internal/app/domain/entity/user"
)

func NewUser(user entity.User) User {
	return User{
		Login:    user.Login,
		Password: user.PasswordHash,
	}
}
