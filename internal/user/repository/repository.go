package repository

import (
	"book-lending-api/internal/user/entity"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewMySQLRepository(db *gorm.DB, log *logrus.Logger) *UserRepository {
	return &UserRepository{db: db, log: log}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) (err error) {
	err = r.db.Create(user).Error
	if err != nil {
		r.log.Errorf("failed to create user: %v", err)
		return
	}
	return
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		r.log.Errorf("failed to find user by email: %v", err)
		return &user, err
	}
	return &user, err
}
