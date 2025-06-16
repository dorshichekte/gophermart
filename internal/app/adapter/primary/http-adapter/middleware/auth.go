package middleware

import (
	"context"
	"gophermarket/internal/constants"
	util2 "gophermarket/internal/util/error_response"
	"net/http"

	"gophermarket/internal/libs/auth"
)

func Auth(auth auth.Auth) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			cookie, cookieErr := req.Cookie(constants.AuthCookieName)
			if cookieErr != nil {
				util2.WriteErrorResponse(res, http.StatusUnauthorized, util2.WrapperError[string]{CustomError: cookieErr.Error()})
				return
			}

			userData, parseAccessTokenErr := auth.ParseAccessToken(cookie.Value)
			if parseAccessTokenErr != nil {
				util2.WriteErrorResponse(res, http.StatusUnauthorized, util2.WrapperError[string]{CustomError: parseAccessTokenErr.Error()})
				return
			}

			ctx := context.WithValue(req.Context(), userIDKey, userData.ID)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}
