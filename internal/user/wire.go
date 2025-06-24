//go:build wireinject
// +build wireinject

package user

import (
	"book-lending-api/config"
	"book-lending-api/internal/middleware"
	"book-lending-api/internal/user/delivery/http"
	"book-lending-api/internal/user/repository"
	"book-lending-api/internal/user/usecase"
	"book-lending-api/pkg/authentication"
	"book-lending-api/pkg/db"
	"book-lending-api/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserHandler() *http.Handler {
	wire.Build(
		repository.NewMySQLRepository,
		wire.Bind(new(repository.IUserRepository), new(*repository.UserRepository)),

		authentication.NewJWTService,
		wire.Bind(new(middleware.IJWTService), new(*authentication.JwtService)),

		usecase.NewUserUseCase,
		wire.Bind(new(usecase.IUserUsecase), new(*usecase.UserUsecase)),

		http.NewHandler,
		ProvideEnvConfig,
		ProvideDatabase,
		logger.NewLogger,
	)
	return &http.Handler{}
}

func ProvideEnvConfig() *config.EnvConfig {
	return config.LoadEnv()
}

func ProvideDatabase(cfg *config.EnvConfig) *gorm.DB {
	return db.InitializeMySQL(cfg)
}
