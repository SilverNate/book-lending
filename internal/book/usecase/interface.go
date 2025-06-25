package usecase

import (
	"book-lending-api/internal/book/dto"
	"book-lending-api/internal/book/entity"
	"context"
)

// mockery --name=IBookUseCase    --dir=internal/book/usecase         --output=internal/book/mocks         --with-expecter
type IBookUseCase interface {
	Create(ctx context.Context, input dto.CreateBookRequest) error
	GetAll(ctx context.Context, offset, limit int) ([]entity.Book, error)
	GetByID(ctx context.Context, id int64) (*entity.Book, error)
	Update(ctx context.Context, id int64, input dto.UpdateBookRequest) error
	Delete(ctx context.Context, id int64) error
}
