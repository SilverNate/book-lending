package usecase

import (
	"book-lending-api/internal/book/dto"
	"book-lending-api/internal/book/entity"
	"context"
)

// mockery --name=IBookUseCase    --dir=internal/book/usecase         --output=internal/book/mocks         --with-expecter
type IBookUseCase interface {
	AddBook(ctx context.Context, input dto.CreateBookRequest) error
	GetAllBooks(ctx context.Context, offset, limit int) ([]entity.Book, error)
	GetBookByID(ctx context.Context, id int64) (*entity.Book, error)
	UpdateBook(ctx context.Context, id int64, input dto.UpdateBookRequest) error
	DeleteBook(ctx context.Context, id int64) error
}
