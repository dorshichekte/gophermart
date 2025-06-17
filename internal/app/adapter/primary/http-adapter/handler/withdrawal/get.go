package withdrawalhandler

import (
	"context"
	"encoding/json"
	"net/http"

	dto "gophermarket/internal/app/adapter/primary/http-adapter/dto/withdrawal"
	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"
	"gophermarket/internal/constants"
	util "gophermarket/internal/util/error_response"
)

func (wh *Handler) GetWithdrawals(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(int)
	if !ok {
		util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: string(constants.ErrFailedGettingUserID)})
		return
	}

	wds, withdrawalErr := wh.Service.Withdrawal.Get(ctx, userID)
	if withdrawalErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: withdrawalErr.Error()})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	withdrawalResponse := make([]dto.WithdrawalResponseDTO, 0, len(wds))
	for _, w := range wds {
		withdrawalResponse = append(withdrawalResponse, dto.NewWithdrawalResponse(w))
	}

	if encodeErr := json.NewEncoder(res).Encode(withdrawalResponse); encodeErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: encodeErr.Error()})
		return
	}
}
