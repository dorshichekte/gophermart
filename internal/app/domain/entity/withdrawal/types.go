package entity

import "time"

type Withdrawal struct {
	OrderNumber string
	Amount      float64
	UserID      int
	CreatedAt   time.Time
}
