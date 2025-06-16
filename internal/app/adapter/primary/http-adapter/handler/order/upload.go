package orderhandler

import (
	"context"
	"io"
	"net/http"
	"strings"

	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"
	"gophermarket/internal/constants"
	util "gophermarket/internal/util/error_response"
)

func (oh *Handler) UploadOrder(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: constants.ErrFailedGettingUserID})
		return
	}

	body, bodyErr := io.ReadAll(req.Body)
	if bodyErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: bodyErr.Error()})
		return
	}

	orderNum := strings.TrimSpace(string(body))
	if orderNum == "" {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: errMissingOrderNumber})
		return
	}

	uploadErr := oh.Service.Order.Upload(ctx, userID, orderNum)
	if uploadErr != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: uploadErr.Error()})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
}
