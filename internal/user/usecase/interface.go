package usecase

import (
	"book-lending-api/internal/user/dto"
	"context"
)

type IUserUsecase interface {
	Register(ctx context.Context, input dto.RegisterRequest) error
	Login(ctx context.Context, input dto.LoginRequest) (string, error)
}
