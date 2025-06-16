package entity

import model "gophermarket/internal/app/repositoriy/model/withdrawal"

func NewWithdrawal(withdrawal model.Withdrawal) Withdrawal {
	return Withdrawal{
		OrderNumber: withdrawal.OrderNumber,
		UserID:      withdrawal.UserID,
		Amount:      withdrawal.Amount,
		CreatedAt:   withdrawal.CreatedAt,
	}
}
