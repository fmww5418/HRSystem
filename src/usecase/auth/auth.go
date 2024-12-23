package auth

import (
	"HRSystem/src/domain/employee"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	dauth "HRSystem/src/domain/auth"
	"HRSystem/src/entity"
	"HRSystem/src/lib/utils"
)

type (
	authUsecase struct {
		repo         dauth.AuthRepository
		employeeRepo employee.EmployeeRepository
		redisClient  *redis.Client
	}
)

var _ dauth.AuthUsecase = (*authUsecase)(nil)

func NewAuthUsecase(repo dauth.AuthRepository, employeeRepo employee.EmployeeRepository, redisClient *redis.Client) dauth.AuthUsecase {
	return &authUsecase{repo: repo, employeeRepo: employeeRepo, redisClient: redisClient}
}

func (s *authUsecase) Register(req dauth.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Username:  req.Username,
		Password:  string(hashedPassword),
		Role:      *req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.WithTransaction(func(tx *gorm.DB) error {
		// Create user within the transaction
		if err := s.repo.CreateUserWithTx(tx, &user); err != nil {
			return err
		}

		// Create employee within the transaction
		if err := s.employeeRepo.CreateWithTx(tx, &entity.Employee{
			ID:          user.ID,
			Name:        user.Username,
			Position:    "Admin",
			ContactInfo: "",
			UserID:      user.ID,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		}); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (s *authUsecase) Login(ctx context.Context, req dauth.LoginRequest) (string, uint, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", 0, errors.New("invalid credentials")
		}
		return "", 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", 0, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(strconv.Itoa(int(user.ID)), user.Role)
	if err != nil {
		return "", 0, err
	}

	// Store session in Redis
	if err := s.redisClient.Set(ctx, dauth.RedisKeyUserSession(user.ID), token, 24*time.Hour).Err(); err != nil {
		// Note: send an event to log center
		fmt.Println("Failed to set token")
	}

	return token, user.ID, nil
}

func (s *authUsecase) Logout(ctx context.Context, userID string) error {
	if err := s.redisClient.Del(ctx, dauth.RedisKeyUserSession(userID)).Err(); err != nil {
		return err
	}

	return nil
}
