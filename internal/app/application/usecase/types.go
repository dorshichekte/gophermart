package usecase

import (
	balance_usecase "gophermarket/internal/app/application/usecase/balance"
	order_usecase "gophermarket/internal/app/application/usecase/order"
	user_usecase "gophermarket/internal/app/application/usecase/user"
	withdrawal_usecase "gophermarket/internal/app/application/usecase/withdrawal"
	balance_repository "gophermarket/internal/app/domain/repository/balance"
	order_repository "gophermarket/internal/app/domain/repository/order"
	user_repository "gophermarket/internal/app/domain/repository/user"
	withdrawal_repository "gophermarket/internal/app/domain/repository/withdrawal"
)

type Repositories struct {
	User       user_repository.UserRepository
	Balance    balance_repository.BalanceRepository
	Order      order_repository.OrderRepository
	Withdrawal withdrawal_repository.WithdrawalRepository
}

type UseCases struct {
	User       *user_usecase.UserUseCase
	Balance    *balance_usecase.BalanceUseCase
	Order      *order_usecase.OrderUseCase
	Withdrawal *withdrawal_usecase.WithdrawalUseCase
}
