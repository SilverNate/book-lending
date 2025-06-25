package usecase

import (
	"book-lending-api/internal/borrow/dto"
	"context"
)

type IBorrowUseCase interface {
	BorrowBooks(ctx context.Context, userID int64, req dto.BorrowRequest) error
	ReturnBooks(ctx context.Context, req dto.ReturnRequest) error
}
