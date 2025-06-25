package usecase

import (
	"book-lending-api/internal/middleware"
	"book-lending-api/internal/user/dto"
	"book-lending-api/internal/user/entity"
	"book-lending-api/internal/user/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUsecase struct {
	repo repository.IUserRepository
	jwt  middleware.IJWTService
}

func NewUserUseCase(repo repository.IUserRepository, jwt middleware.IJWTService) *UserUsecase {
	return &UserUsecase{repo, jwt}
}

func (u *UserUsecase) Register(ctx context.Context, input dto.RegisterRequest) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := entity.User{
		Email:        input.Email,
		PasswordHash: string(hashed),
		CreatedAt:    time.Now(),
	}
	return u.repo.CreateUser(ctx, &user)
}

func (u *UserUsecase) Login(ctx context.Context, input dto.LoginRequest) (string, error) {
	user, err := u.repo.FindUserByEmail(ctx, input.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) != nil {
		return "", err
	}
	return u.jwt.GenerateToken(user.ID, user.Email)
}
