package model

import "time"

type Balance struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Current   float64   `db:"current"`
	Withdrawn float64   `db:"withdrawn"`
	CreatedAt time.Time `db:"created_at"`
}
