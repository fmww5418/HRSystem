package auth

import (
	"gorm.io/gorm"
	"time"

	dauth "HRSystem/src/domain/auth"
	"HRSystem/src/entity"
)

type authRepository struct {
	db *gorm.DB
}

var _ dauth.AuthRepository = (*authRepository)(nil)

func NewAuthRepository(db *gorm.DB) dauth.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(user entity.User) error {
	return r.db.Create(&user).Error
}

func (r *authRepository) UpdateUser(user entity.User) error {
	user.UpdatedAt = time.Now()
	return r.db.Save(&user).Error
}

func (r *authRepository) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *authRepository) FindByUserID(userID uint) (entity.User, error) {
	var user entity.User
	err := r.db.Where("ID = ?", userID).First(&user).Error
	return user, err
}

func (r *authRepository) CreateUserWithTx(tx *gorm.DB, user *entity.User) error {
	return tx.Create(user).Error
}

func (r *authRepository) WithTransaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}
