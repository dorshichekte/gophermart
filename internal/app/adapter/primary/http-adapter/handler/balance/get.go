package balancehandler

import (
	"context"
	"encoding/json"
	"net/http"

	dto "gophermarket/internal/app/adapter/primary/http-adapter/dto/balance"
	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"
	"gophermarket/internal/constants"
	util "gophermarket/internal/util/error_response"
)

func (bh *Handler) GetBalance(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(int)
	if !ok {
		util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: string(constants.ErrFailedGettingUserID)})
		return
	}

	balance, serviceErr := bh.Service.Balance.Get(ctx, userID)
	if serviceErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: serviceErr.Error()})
		return
	}

	respBalance := dto.NewBalanceResponse(balance.Current, balance.Withdrawn)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(res).Encode(respBalance); encodeErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: encodeErr.Error()})
	}
}
