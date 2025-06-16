package model

import (
	"time"
)

type User struct {
	ID         int       `db:"id"`
	Login      string    `db:"login"`
	Password   string    `db:"password"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
	IsActive   bool      `db:"is_active"`
}
