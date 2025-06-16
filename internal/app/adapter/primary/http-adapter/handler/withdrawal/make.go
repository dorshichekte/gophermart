package withdrawalhandler

import (
	"context"
	"net/http"

	dto "gophermarket/internal/app/adapter/primary/http-adapter/dto/withdrawal"
	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"
	"gophermarket/internal/constants"
	v "gophermarket/internal/libs/validator"
	util "gophermarket/internal/util/error_response"
)

func (wh *Handler) Make(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeRequest)
	defer cancel()
	defer func() {
		_ = req.Body.Close()
	}()

	userID, ok := req.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: constants.ErrFailedGettingUserID})
		return
	}

	var withdrawalDto dto.WithdrawalRequestDTO
	decodeErr := wh.DecodeJSON(req, &withdrawalDto)
	if decodeErr != nil {
		util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[string]{CustomError: decodeErr.Error()})
		return
	}

	validateErr := wh.validator.ValidateStruct(&withdrawalDto)
	if validateErr != nil {
		validationErrors, parseError := wh.validator.ParseValidationErrors(validateErr)
		if parseError != nil {
			util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[string]{CustomError: validateErr.Error()})
			return
		}

		util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[[]v.ValidationError]{CustomError: validationErrors})
		return
	}

	makeErr := wh.Service.Withdrawal.Make(ctx, withdrawalDto.Order, userID, withdrawalDto.Sum)
	if makeErr != nil {
		util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[string]{CustomError: makeErr.Error()})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
