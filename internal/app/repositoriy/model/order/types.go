package model

import (
	"time"

	entity "gophermarket/internal/app/domain/entity/order"
)

type Order struct {
	ID         int           `db:"id"`
	Number     string        `db:"number"`
	Status     entity.Status `db:"status"`
	UserID     int           `db:"user_id"`
	Accrual    float64       `db:"accrual"`
	Active     bool          `db:"active"`
	UploadAt   time.Time     `db:"upload_at"`
	ModifiedAt time.Time     `db:"modified_at"`
}
