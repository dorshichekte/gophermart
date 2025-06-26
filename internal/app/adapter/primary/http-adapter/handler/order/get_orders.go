package orderhandler

import (
	"context"
	"encoding/json"
	"net/http"

	dto "gophermarket/internal/app/adapter/primary/http-adapter/dto/order"
	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"
	"gophermarket/internal/constants"
	util "gophermarket/internal/util/error_response"
)

func (oh *Handler) GetOrders(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(int)
	if !ok {
		util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: constants.ErrFailedGettingUserID.Error()})
		return
	}

	orders, getOrdersErr := oh.ServiceOrder.GetAll(ctx, userID)
	if getOrdersErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: getOrdersErr.Error()})
		return
	}

	isOrdersEmpty := len(orders) == 0
	if isOrdersEmpty {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNoContent)
		return
	}

	ordersResponse := make([]dto.OrderResponseDTO, 0, len(orders))
	for _, order := range orders {
		orderResponse := dto.NewOrderResponse(order.Number, order.Status, order.Accrual, order.UploadedAt)
		ordersResponse = append(ordersResponse, orderResponse)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(res).Encode(ordersResponse); encodeErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: encodeErr.Error()})
		return
	}
}
