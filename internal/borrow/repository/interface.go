package repository

import (
	"book-lending-api/internal/borrow/entity"
	"context"
)

// mockery --name=IBorrowRepository    --dir=internal/borrow/repository         --output=internal/borrow/mocks         --with-expecter
type IBorrowRepository interface {
	CreateBorrowing(ctx context.Context, b *entity.Borrowing) error
	FindBorrowingByID(ctx context.Context, id int64) (*entity.Borrowing, error)
	FindActiveByUserAndBook(ctx context.Context, userID, bookID int64) (*entity.Borrowing, error)
	InsertOrUpdateBorrowing(ctx context.Context, b *entity.Borrowing) error
	CountBorrowsInLast7Days(ctx context.Context, userID int64) (int, error)
}
