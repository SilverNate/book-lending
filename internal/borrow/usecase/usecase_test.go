package usecase_test

import (
	"book-lending-api/internal/book/entity"
	mockBook "book-lending-api/internal/book/mocks"
	"book-lending-api/internal/borrow/dto"
	modelBorrow "book-lending-api/internal/borrow/entity"
	"book-lending-api/internal/borrow/mocks"
	"book-lending-api/internal/borrow/usecase"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBorrow_Success(t *testing.T) {
	bookRepo := new(mockBook.IBookUseCase)
	borrowRepo := new(mocks.IBorrowRepository)
	log := logrus.New()
	uc := usecase.NewBorrowUsecase(borrowRepo, log, bookRepo)

	bookRepo.On("GetBookByID", mock.Anything, int64(10)).Return(&entity.Book{ID: 10, Quantity: 3}, nil)
	borrowRepo.On("CountBorrowsInLast7Days", mock.Anything, int64(1)).Return(1, nil)
	borrowRepo.On("FindActiveByUserAndBook", mock.Anything, int64(1), int64(10)).Return(nil, errors.New("not found"))
	borrowRepo.On("CreateBorrowing", mock.Anything, mock.AnythingOfType("*entity.Borrowing")).Return(nil)

	err := uc.BorrowBooks(context.Background(), 1, dto.BorrowRequest{BookID: 10})
	assert.NoError(t, err)
}

func TestReturn_Success(t *testing.T) {
	bookRepo := new(mockBook.IBookUseCase)
	borrowRepo := new(mocks.IBorrowRepository)
	log := logrus.New()
	uc := usecase.NewBorrowUsecase(borrowRepo, log, bookRepo)

	now := time.Now()
	borrow := &modelBorrow.Borrowing{
		ID:         1,
		BookID:     10,
		UserID:     1,
		BorrowDate: now,
	}

	borrowRepo.On("FindBorrowingByID", mock.Anything, int64(1)).Return(borrow, nil)
	borrowRepo.On("InsertOrUpdateBorrowing", mock.Anything, mock.MatchedBy(func(b *modelBorrow.Borrowing) bool {
		return b.ReturnDate != nil
	})).Return(nil)

	err := uc.ReturnBooks(context.Background(), dto.ReturnRequest{BorrowingID: 1})
	assert.NoError(t, err)
}

func TestBorrow_Fail_BookNotFound(t *testing.T) {
	bookRepo := new(mockBook.IBookUseCase)
	borrowRepo := new(mocks.IBorrowRepository)
	log := logrus.New()
	uc := usecase.NewBorrowUsecase(borrowRepo, log, bookRepo)

	bookRepo.On("GetBookByID", mock.Anything, int64(99)).Return(nil, errors.New("not found"))

	err := uc.BorrowBooks(context.Background(), 1, dto.BorrowRequest{BookID: 99})
	assert.EqualError(t, err, "book not found")
}
