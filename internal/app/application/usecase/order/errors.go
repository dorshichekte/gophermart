package orderusecase

import customerror "gophermarket/internal/error"

var (
	ErrOrderExists              = customerror.New(OrderExists)
	ErrOrderExistsByAnotherUser = customerror.New(OrderExistsByAnotherUser)
)
