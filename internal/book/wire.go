//go:build wireinject
// +build wireinject

package book

import (
	"book-lending-api/config"
	"book-lending-api/internal/book/delivery/http"
	"book-lending-api/internal/book/repository"
	"book-lending-api/internal/book/usecase"
	"book-lending-api/pkg/db"
	logger "book-lending-api/pkg/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitBookHandler() *http.Handler {
	wire.Build(
		repository.NewBookRepository,
		wire.Bind(new(repository.IBookRepository), new(*repository.BookRepository)),

		usecase.NewBookUseCase,
		wire.Bind(new(usecase.IBookUseCase), new(*usecase.BookUsecase)),
		http.NewHandler,

		ProvideEnvConfig,
		ProvideDatabase,
		ProvideLogger,
	)
	return &http.Handler{}
}

func ProvideEnvConfig() *config.EnvConfig {
	return config.LoadEnv()
}

func ProvideDatabase(cfg *config.EnvConfig) *gorm.DB {
	return db.InitializeMySQL(cfg)
}

func ProvideLogger() *logrus.Logger {
	return logger.NewLogger()
}
