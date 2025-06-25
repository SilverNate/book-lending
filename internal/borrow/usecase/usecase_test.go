package usecase_test

import (
	"book-lending-api/internal/book/entity"
	mockBook "book-lending-api/internal/book/mocks"
	"book-lending-api/internal/borrow/dto"
	modelBorrow "book-lending-api/internal/borrow/entity"
	"book-lending-api/internal/borrow/mocks"
	"book-lending-api/internal/borrow/usecase"
	"book-lending-api/pkg/infrastructure"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	redismock "github.com/go-redis/redismock/v8"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRedisWrapper struct {
	Client *redis.Client
}

func setup() (*mockBook.IBookUseCase, *mocks.IBorrowRepository, redismock.ClientMock, *usecase.BorrowUsecase) {
	bookUsecase := new(mockBook.IBookUseCase)
	borrowRepo := new(mocks.IBorrowRepository)
	log := logrus.New()

	rdb, redisMock := redismock.NewClientMock()
	mockRedis := &infrastructure.RedisClient{Client: rdb}

	uc := usecase.NewBorrowUsecase(borrowRepo, log, bookUsecase, mockRedis)
	return bookUsecase, borrowRepo, redisMock, uc
}

func mockRedisSuccess(mockRedis redismock.ClientMock) {
	mockRedis.ExpectGet("some-key").SetVal("0")
	mockRedis.ExpectIncr("some-key").SetVal(1)
	mockRedis.ExpectExpire("some-key", time.Hour).SetVal(true)
	mockRedis.ExpectTxPipeline()
	mockRedis.ExpectTxPipelineExec()
}

func TestBorrow_Success(t *testing.T) {
	bookUsecase, borrowRepo, redisMock, uc := setup()
	ctx := context.Background()

	mockRedisSuccess(redisMock)

	bookUsecase.On("GetBookByID", mock.Anything, mock.Anything).Return(&entity.Book{ID: 10, Quantity: 3}, nil)
	borrowRepo.On("IsBookBorrowed", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	borrowRepo.On("CreateBorrowing", mock.Anything, mock.Anything).Return(nil)
	bookUsecase.On("UpdateBook", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := uc.BorrowBooks(ctx, 1, dto.BorrowRequest{BookID: 10})
	assert.NoError(t, err)
}

func TestBorrow_BookNotFound(t *testing.T) {
	bookUsecase, _, redisMock, uc := setup()
	ctx := context.Background()

	mockRedisSuccess(redisMock)
	bookUsecase.On("GetBookByID", mock.Anything, int64(999)).Return(nil, errors.New("book not found"))

	err := uc.BorrowBooks(ctx, 1, dto.BorrowRequest{BookID: 999})
	assert.Error(t, err)
	assert.EqualError(t, err, "book not found")
}

func TestBorrow_ZeroQuantity(t *testing.T) {
	bookUsecase, _, redisMock, uc := setup()
	ctx := context.Background()

	mockRedisSuccess(redisMock)

	bookUsecase.On("GetBookByID", mock.Anything, int64(10)).Return(&entity.Book{ID: 10, Quantity: 0}, nil)

	err := uc.BorrowBooks(ctx, 1, dto.BorrowRequest{BookID: 10})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not available")
}

func TestBorrow_AlreadyBorrowed(t *testing.T) {
	bookUsecase, borrowRepo, redisMock, uc := setup()
	ctx := context.Background()

	mockRedisSuccess(redisMock)

	bookUsecase.On("GetBookByID", mock.Anything, mock.Anything).Return(&entity.Book{ID: 10, Quantity: 3}, nil)
	borrowRepo.On("IsBookBorrowed", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)

	err := uc.BorrowBooks(ctx, 1, dto.BorrowRequest{BookID: 10})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already borrowed")
}

func TestBorrow_CreateBorrowingFails(t *testing.T) {
	bookUsecase, borrowRepo, redisMock, uc := setup()
	ctx := context.Background()

	mockRedisSuccess(redisMock)

	bookUsecase.On("GetBookByID", mock.Anything, mock.Anything).Return(&entity.Book{ID: 10, Quantity: 3}, nil)
	borrowRepo.On("IsBookBorrowed", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	borrowRepo.On("CreateBorrowing", mock.Anything, mock.Anything).Return(errors.New("error create borrowing"))

	err := uc.BorrowBooks(ctx, 1, dto.BorrowRequest{BookID: 10})
	assert.Error(t, err)
	assert.EqualError(t, err, "error create borrowing")
}

func TestBorrow_RedisFail(t *testing.T) {
	bookUsecase, borrowRepo, redisMock, uc := setup()
	ctx := context.Background()

	redisMock.ExpectGet("borrow-limit:1").SetVal("10")

	bookUsecase.On("GetBookByID", mock.Anything, mock.Anything).Return(&entity.Book{ID: 10, Quantity: 3}, nil)
	borrowRepo.On("IsBookBorrowed", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	borrowRepo.On("CreateBorrowing", mock.Anything, mock.Anything).Return(nil)
	bookUsecase.On("UpdateBook", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := uc.BorrowBooks(ctx, 1, dto.BorrowRequest{BookID: 10})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rate limit exceeded: max 5 borrows per 7 days")

	// Ensure all expectations were met
	assert.NoError(t, redisMock.ExpectationsWereMet())

}

func TestReturn_Success(t *testing.T) {
	bookUsecase, borrowRepo, _, uc := setup()
	ctx := context.Background()

	now := time.Now()
	borrow := &modelBorrow.Borrowing{
		ID:         1,
		BookID:     10,
		UserID:     1,
		BorrowDate: now,
	}

	borrowRepo.On("GetBorrowingByID", mock.Anything, mock.Anything).
		Return(borrow, nil)

	bookUsecase.On("GetBookByID", mock.Anything, mock.Anything).Return(&entity.Book{ID: 10, Quantity: 3}, nil)
	bookUsecase.On("UpdateBook", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	borrowRepo.On("InsertOrUpdateBorrowing", mock.Anything, mock.MatchedBy(func(b *modelBorrow.Borrowing) bool {
		return b.ReturnDate != nil
	})).Return(nil)

	err := uc.ReturnBooks(ctx, dto.ReturnRequest{BorrowingID: 1})
	assert.NoError(t, err)
}

func TestReturn_Fail_BorrowingNotFound(t *testing.T) {
	_, borrowRepo, _, uc := setup()
	ctx := context.Background()

	borrowRepo.On("GetBorrowingByID", mock.Anything, mock.Anything).
		Return(nil, errors.New("borrowing record not found"))

	err := uc.ReturnBooks(ctx, dto.ReturnRequest{BorrowingID: 2})
	assert.Error(t, err)
	assert.EqualError(t, err, "borrowing record not found")
}

func TestReturn_Fail_UpdateError(t *testing.T) {
	_, borrowRepo, _, uc := setup()
	ctx := context.Background()

	now := time.Now()
	borrow := &modelBorrow.Borrowing{
		ID:         3,
		BookID:     11,
		UserID:     1,
		BorrowDate: now,
	}

	borrowRepo.On("GetBorrowingByID", mock.Anything, mock.Anything).
		Return(borrow, nil)

	borrowRepo.On("InsertOrUpdateBorrowing", mock.Anything, mock.AnythingOfType("*entity.Borrowing")).
		Return(errors.New("update failed"))

	err := uc.ReturnBooks(ctx, dto.ReturnRequest{BorrowingID: 3})
	assert.Error(t, err)
	assert.EqualError(t, err, "update failed")
}
