//go:build wireinject
// +build wireinject

package borrow

import (
	"book-lending-api/config"
	bookRepo "book-lending-api/internal/book/repository"
	bookUsecase "book-lending-api/internal/book/usecase"
	httpDelivery "book-lending-api/internal/borrow/delivery/http"
	repo "book-lending-api/internal/borrow/repository"
	"book-lending-api/internal/borrow/usecase"
	"book-lending-api/pkg/db"
	"book-lending-api/pkg/infrastructure"
	"book-lending-api/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitBorrowHandler() *httpDelivery.BorrowHandler {
	wire.Build(
		repo.NewBorrowRepository,
		wire.Bind(new(repo.IBorrowRepository), new(*repo.BorrowRepository)),

		usecase.NewBorrowUsecase,
		wire.Bind(new(usecase.IBorrowUseCase), new(*usecase.BorrowUsecase)),

		bookUsecase.NewBookUseCase,
		wire.Bind(new(bookUsecase.IBookUseCase), new(*bookUsecase.BookUsecase)),

		bookRepo.NewBookRepository,
		wire.Bind(new(bookRepo.IBookRepository), new(*bookRepo.BookRepository)),

		httpDelivery.NewBorrowHandler,
		ProvideEnvConfig,
		ProvideDatabase,
		logger.NewLogger,
		ProvideRedis,
	)
	return &httpDelivery.BorrowHandler{}
}

func ProvideEnvConfig() *config.EnvConfig {
	return config.LoadEnv()
}

func ProvideDatabase(cfg *config.EnvConfig) *gorm.DB {
	return db.InitializeMySQL(cfg)
}

func ProvideRedis() *infrastructure.RedisClient {
	return infrastructure.NewRedis()
}
