//go:build unit

package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"HRSystem/src/domain/auth"
	dauth "HRSystem/src/domain/auth"
	"HRSystem/src/domain/auth/mocks"
	mEmployee "HRSystem/src/domain/employee/mocks"
	"HRSystem/src/entity"
	rlib "HRSystem/src/lib/redis"
)

func TestLoginSuccess(t *testing.T) {
	fakePassword := "password123"
	fakeHashedPassword, _ := bcrypt.GenerateFromPassword([]byte(fakePassword), bcrypt.DefaultCost)
	type (
		testCase struct {
			name        string
			user        entity.User
			reqPassword string
			mockFn      func(
				t *testing.T,
				user entity.User,
				exError error,
			) dauth.AuthUsecase
			expectUserID uint
			expectErr    error
		}
	)

	var (
		testCases = []testCase{
			{
				name: "Successfully login",
				user: entity.User{
					ID:       1,
					Username: "username",
					Password: string(fakeHashedPassword),
					Role:     entity.RoleAdmin,
				},
				reqPassword: fakePassword,
				mockFn: func(
					t *testing.T,
					user entity.User,
					exError error,
				) dauth.AuthUsecase {
					mockRepo := mocks.NewMockAuthRepository(t)
					employeeMockRepo := mEmployee.NewMockEmployeeRepository(t)
					usecase := NewAuthUsecase(mockRepo, employeeMockRepo, rlib.DB.Client)

					mockRepo.EXPECT().FindByUsername(user.Username).Return(user, nil).Once()

					return usecase
				},
				expectUserID: 1,
				expectErr:    nil,
			},
			{
				name: "failed to login",
				user: entity.User{
					ID:       1,
					Username: "username",
					Password: string(fakeHashedPassword),
					Role:     entity.RoleAdmin,
				},
				reqPassword: "invalid_password",
				mockFn: func(
					t *testing.T,
					user entity.User,
					exError error,
				) dauth.AuthUsecase {
					mockRepo := mocks.NewMockAuthRepository(t)
					employeeMockRepo := mEmployee.NewMockEmployeeRepository(t)
					usecase := NewAuthUsecase(mockRepo, employeeMockRepo, rlib.DB.Client)

					mockRepo.EXPECT().FindByUsername(user.Username).Return(user, nil).Once()

					return usecase
				},
				expectUserID: 0,
				expectErr:    errors.New("invalid credentials"),
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			uLogin := tc.mockFn(tt, tc.user, nil)

			token, returnedUserID, err := uLogin.Login(context.Background(), auth.LoginRequest{
				Username: tc.user.Username,
				Password: tc.reqPassword,
			})

			assert.Equal(t, tc.expectUserID, returnedUserID)
			assert.Equal(tt, tc.expectErr, err)
			if err == nil {
				assert.NotEmpty(t, token)
			}
		})
	}
}
