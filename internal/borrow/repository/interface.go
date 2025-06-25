package repository

import (
	"book-lending-api/internal/borrow/entity"
	"context"
)

// mockery --name=IBorrowRepository    --dir=internal/borrow/repository         --output=internal/borrow/mocks         --with-expecter
type IBorrowRepository interface {
	CreateBorrowing(ctx context.Context, b *entity.Borrowing) error
	GetBorrowingByID(ctx context.Context, id int64) (*entity.Borrowing, error)
	IsBookBorrowed(ctx context.Context, userID, bookID int64) (bool, error)
	InsertOrUpdateBorrowing(ctx context.Context, b *entity.Borrowing) error
}
