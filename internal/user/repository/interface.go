package repository

import (
	"book-lending-api/internal/user/entity"
	"context"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
