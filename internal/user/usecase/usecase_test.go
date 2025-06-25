package usecase_test

import (
	jwtMocks "book-lending-api/internal/middleware/mocks"
	"book-lending-api/internal/user/dto"
	"book-lending-api/internal/user/entity"
	"book-lending-api/internal/user/mocks"
	"book-lending-api/internal/user/usecase"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.IUserRepository)
	mockJWT := new(jwtMocks.IJWTService)

	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	uc := usecase.NewUserUseCase(mockRepo, mockJWT)
	err := uc.Register(context.Background(), dto.RegisterRequest{
		Email:    "test@mail.com",
		Password: "test1234",
	})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	mockRepo.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mocks.IUserRepository)
	mockJWT := new(jwtMocks.IJWTService)

	hashed := "$2a$10$UF1GV7frKJC5oqXwHWERzOvGR/ftNfBDMvSbxhKAlQ6q/cltS00xK" // bcrypt("secret123")
	mockRepo.On("FindUserByEmail", mock.Anything, "test@mail.com").Return(&entity.User{
		ID:           1,
		Email:        "test@mail.com",
		PasswordHash: hashed,
	}, nil)

	mockJWT.On("GenerateToken", mock.AnythingOfType("int64"), "test@mail.com").Return("mock.token", nil)

	uc := usecase.NewUserUseCase(mockRepo, mockJWT)
	token, err := uc.Login(context.Background(), dto.LoginRequest{
		Email:    "test@mail.com",
		Password: "secret123",
	})

	t.Log("token:", token)

	assert.NoError(t, err)
	assert.Equal(t, "mock.token", token)

	mockRepo.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}
