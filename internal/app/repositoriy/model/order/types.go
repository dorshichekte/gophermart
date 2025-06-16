package model

import (
	entity "gophermarket/internal/app/domain/entity/order"
	"time"
)

type Order struct {
	ID         int           `db:"id"`
	Number     string        `db:"number"`
	Status     entity.Status `db:"status"`
	UserID     int           `db:"user_id"`
	UploadedAt time.Time     `db:"uploaded_at"`
	ModifiedAt time.Time     `db:"modified_at"`
	Active     bool          `db:"active"`
	Accrual    float64       `db:"accrual"`
}
