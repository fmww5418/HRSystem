package auth

import (
	"context"
	"gorm.io/gorm"

	"HRSystem/src/entity"
)

type (
	AuthRepository interface {
		CreateUser(user entity.User) error
		UpdateUser(user entity.User) error
		FindByUsername(username string) (entity.User, error)
		FindByUserID(userID uint) (entity.User, error)
		CreateUserWithTx(tx *gorm.DB, user *entity.User) error
		WithTransaction(fn func(tx *gorm.DB) error) error
	}

	AuthUsecase interface {
		Register(req RegisterRequest) error
		Login(ctx context.Context, req LoginRequest) (string, uint, error)
		Logout(ctx context.Context, userID string) error
	}
)

type (
	RegisterRequest struct {
		Username string       `json:"username" binding:"required"`
		Password string       `json:"password" binding:"required"`
		Role     *entity.Role `json:"role" binding:"required"`
	}
	LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)
