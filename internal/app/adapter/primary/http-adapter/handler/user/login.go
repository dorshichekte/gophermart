package userhandler

import (
	"context"
	"errors"
	"net/http"

	"gophermarket/internal/app/adapter/primary/http-adapter/dto/user"
	user_usecase "gophermarket/internal/app/application/usecase/user"
	constants2 "gophermarket/internal/constants"
	"gophermarket/internal/libs/auth"
	v "gophermarket/internal/libs/validator"
	util "gophermarket/internal/util/error_response"
)

func (uh *Handler) Login(auth auth.Auth) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), constants2.DefaultTimeRequest)
		defer cancel()
		defer func() {
			_ = req.Body.Close()
		}()

		var loginRequestDto dto.LoginRequest
		decodeErr := uh.DecodeJSON(req, &loginRequestDto)
		if decodeErr != nil {
			uh.Logger.Error(decodeErr.Error())
			util.WriteErrorResponse(res, http.StatusBadRequest, util.WrapperError[string]{CustomError: decodeErr.Error()})
			return
		}

		validateErr := uh.validator.ValidateStruct(&loginRequestDto)
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

		userID, loginErr := uh.Service.User.Login(ctx, loginRequestDto.Login, loginRequestDto.Password)
		if loginErr != nil {
			if errors.Is(loginErr, user_usecase.ErrUserNotFound) {
				uh.Logger.Error(loginErr.Error())
				util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: loginErr.Error()})
				return
			}

			uh.Logger.Error(loginErr.Error())
			util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: loginErr.Error()})
			return
		}

		authData, setAuthErr := auth.Generate(userID)
		if setAuthErr != nil {
			uh.Logger.Error(setAuthErr.Error())
			util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: setAuthErr.Error()})
			return
		}

		http.SetCookie(res, &http.Cookie{
			Name:  constants2.AuthCookieName,
			Value: authData.AccessToken,
			Path:  "/",
		})

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
	}
}
