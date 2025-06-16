package model

import (
	"time"
)

type Withdrawal struct {
	ID          int       `db:"id"`
	OrderNumber string    `db:"order"`
	Amount      float64   `db:"sum"`
	UserID      int       `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
}
