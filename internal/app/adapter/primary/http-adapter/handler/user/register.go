package userhandler

import (
	"context"
	"errors"
	"net/http"

	"gophermarket/internal/app/adapter/primary/http-adapter/dto/user"
	user_usecase "gophermarket/internal/app/application/usecase/user"
	"gophermarket/internal/constants"
	"gophermarket/internal/libs/auth"
	v "gophermarket/internal/libs/validator"
	util "gophermarket/internal/util/error_response"
)

func (uh *Handler) Register(auth auth.Auth) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeRequest)
		defer cancel()
		defer func() {
			_ = req.Body.Close()
		}()

		var registerRequestDto dto.RegisterRequest
		decodeErr := uh.DecodeJSON(req, &registerRequestDto)
		if decodeErr != nil {
			uh.Logger.Error(decodeErr.Error())
			util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[string]{CustomError: decodeErr.Error()})
			return
		}

		validateErr := uh.validator.ValidateStruct(&registerRequestDto)
		if validateErr != nil {
			validationErrors, parseError := uh.validator.ParseValidationErrors(validateErr)
			if parseError != nil {
				uh.Logger.Error(parseError.Error())
				util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[string]{CustomError: validateErr.Error()})
				return
			}

			uh.Logger.Error(validateErr.Error())
			util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[[]v.ValidationError]{CustomError: validationErrors})
			return
		}

		userID, registerErr := uh.Service.User.Register(ctx, registerRequestDto.Login, registerRequestDto.Password)
		if registerErr != nil {
			uh.Logger.Error(registerErr.Error())
			if errors.Is(registerErr, user_usecase.ErrLoginAlreadyTaken) {
				util.WriteErrorResponse(res, http.StatusConflict, util.WrapperError[string]{CustomError: registerErr.Error()})
				return
			}

			util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: registerErr.Error()})
			return
		}

		authData, setAuthErr := auth.Generate(userID)
		if setAuthErr != nil {
			uh.Logger.Error(setAuthErr.Error())
			util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: setAuthErr.Error()})
			return
		}

		http.SetCookie(res, &http.Cookie{
			Name:  constants.AuthCookieName,
			Value: authData.AccessToken,
			Path:  "/",
		})

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
	}
}
