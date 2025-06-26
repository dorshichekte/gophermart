package userhandler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	dto "gophermarket/internal/app/adapter/primary/http-adapter/dto/user"
	base_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/base"
	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"
	user_usecase "gophermarket/internal/app/application/usecase/user"
	"gophermarket/internal/app/config"
	env "gophermarket/internal/app/config/env"
	entity "gophermarket/internal/app/domain/entity/user"
	mocks "gophermarket/internal/app/domain/mock"
	model "gophermarket/internal/app/repositoriy/model/user"
	"gophermarket/internal/libs/auth"
	"gophermarket/internal/libs/hasher"
	"gophermarket/internal/libs/logger"
	v "gophermarket/internal/libs/validator"
)

func TestRegisterUserHandler(t *testing.T) {
	cfg := config.Config{Env: env.Env{ServerAddress: "localhost:8080"}}

	l, loggerErr := logger.New()
	require.NoError(t, loggerErr)

	bh := base_handler.New(l)
	h := hasher.New()
	validator := v.New()
	a := auth.New(cfg.Env.AccessSecretKey)
	login := fmt.Sprintf("test%d", time.Now().UnixNano())[:10]

	testCases := []struct {
		name         string
		body         dto.RegisterRequest
		expectedCode int
		buildStubs   func(s *mocks.MockUserRepository)
		userID       int
	}{
		{
			name: "Register user",
			body: dto.RegisterRequest{
				AuthRequest: dto.AuthRequest{
					Login:    login,
					Password: "12345678",
				},
			},
			expectedCode: http.StatusOK,
			userID:       1,
			buildStubs: func(r *mocks.MockUserRepository) {
				r.EXPECT().
					GetByLogin(gomock.Any(), login).
					Return(&model.User{}, sql.ErrNoRows).
					Times(1)

				r.EXPECT().
					Register(gomock.Any(), gomock.AssignableToTypeOf(entity.User{
						Login:        login,
						PasswordHash: gomock.Any().String(),
					})).
					Times(1).
					Return(1, nil)
			},
		},
		{
			name: "Register user with empty fields (login, password)",
			body: dto.RegisterRequest{
				AuthRequest: dto.AuthRequest{
					Login:    "",
					Password: "",
				},
			},
			expectedCode: http.StatusBadRequest,
			userID:       1,
			buildStubs: func(r *mocks.MockUserRepository) {
				r.EXPECT().
					Register(gomock.Any(), gomock.Any()).
					Times(0)
			},
		},
		{
			name: "Login too short",
			body: dto.RegisterRequest{
				AuthRequest: dto.AuthRequest{
					Login:    "te",
					Password: "12345678",
				},
			},
			expectedCode: http.StatusBadRequest,
			userID:       1,
			buildStubs: func(r *mocks.MockUserRepository) {
				r.EXPECT().
					Register(gomock.Any(), gomock.Any()).
					Times(0)
			},
		},
		{
			name: "Password too short",
			body: dto.RegisterRequest{
				AuthRequest: dto.AuthRequest{
					Login:    login,
					Password: "12",
				},
			},
			expectedCode: http.StatusBadRequest,
			userID:       1,
			buildStubs: func(r *mocks.MockUserRepository) {
				r.EXPECT().
					Register(gomock.Any(), gomock.Any()).
					Times(0)
			},
		},
		{
			name: "User already registered",
			body: dto.RegisterRequest{
				AuthRequest: dto.AuthRequest{
					Login:    login,
					Password: "12345678",
				},
			},
			expectedCode: http.StatusConflict,
			userID:       1,
			buildStubs: func(r *mocks.MockUserRepository) {
				r.EXPECT().
					GetByLogin(gomock.Any(), login).
					Return(&model.User{}, user_usecase.ErrLoginAlreadyTaken).
					Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			json, err := json.Marshal(tc.body)
			require.NoError(t, err)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockUserRepository(ctrl)
			tc.buildStubs(repo)

			req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(json))
			ctx := context.WithValue(req.Context(), middleware.UserIDKey(), tc.userID)
			req = req.WithContext(ctx)

			w := httptest.NewRecorder()
			uc := user_usecase.New(l, h, a, repo)

			handler := New(bh, l, uc, validator)

			handler.Register(a)(w, req)

			require.Equal(t, tc.expectedCode, w.Code)

			if tc.expectedCode == http.StatusOK {
				require.Equal(t, "application/json", w.Header().Get("Content-Type"))
			}
		})
	}
}
