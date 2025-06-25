package usecase

import (
	dtoBook "book-lending-api/internal/book/dto"
	modelBook "book-lending-api/internal/book/entity"
	"book-lending-api/internal/book/usecase"
	"book-lending-api/internal/borrow/dto"
	"book-lending-api/internal/borrow/entity"
	"book-lending-api/internal/borrow/repository"
	"book-lending-api/pkg/infrastructure"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type BorrowUsecase struct {
	repo        repository.IBorrowRepository
	log         *logrus.Logger
	bookUsecase usecase.IBookUseCase
	redis       *infrastructure.RedisClient
}

func NewBorrowUsecase(repo repository.IBorrowRepository, log *logrus.Logger, bookUsecase usecase.IBookUseCase, redis *infrastructure.RedisClient) *BorrowUsecase {
	return &BorrowUsecase{repo: repo, log: log, bookUsecase: bookUsecase, redis: redis}
}

func (uc *BorrowUsecase) BorrowBooks(ctx context.Context, userID int64, req dto.BorrowRequest) error {
	uc.log.Infof("borrowing books userId: %v, request: %v", userID, req)

	if uc.isRateLimited(ctx, userID) {
		return errors.New("rate limit exceeded: max 5 borrows per 7 days")
	}

	book, err := uc.bookUsecase.GetBookByID(ctx, req.BookID)
	if err != nil || book == nil {
		uc.log.Errorf("error get book when borrowing: %v", err)
		return errors.New("book not found")
	}

	if book.Quantity < 1 {
		return errors.New("book not available")
	}

	isBorrowed, err := uc.repo.IsBookBorrowed(ctx, userID, book.ID)
	if err != nil {
		return err
	}
	if isBorrowed {
		return errors.New("book is already borrowed")
	}

	borrow := &entity.Borrowing{
		BookID:     req.BookID,
		UserID:     userID,
		BorrowDate: time.Now(),
	}

	err = uc.repo.CreateBorrowing(ctx, borrow)
	if err != nil {
		return err
	}

	updateBook := mappingUpdateBook(book, book.Quantity-1)
	uc.log.Debug("borrow book: ", updateBook)

	err = uc.bookUsecase.UpdateBook(ctx, book.ID, updateBook)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BorrowUsecase) ReturnBooks(ctx context.Context, req dto.ReturnRequest) error {
	uc.log.Infof("return books request: %v", req)

	borrow, err := uc.repo.GetBorrowingByID(ctx, req.BorrowingID)
	if err != nil {
		return errors.New("borrowing record not found")
	}
	now := time.Now()
	borrow.ReturnDate = &now

	if err := uc.repo.InsertOrUpdateBorrowing(ctx, borrow); err != nil {
		return err
	}

	book, err := uc.bookUsecase.GetBookByID(ctx, borrow.BookID)
	if err != nil || book == nil {
		return errors.New("book not found during return")
	}

	updateBook := mappingUpdateBook(book, book.Quantity+1)
	uc.log.Debug("return  book: ", updateBook)

	err = uc.bookUsecase.UpdateBook(ctx, book.ID, updateBook)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BorrowUsecase) isRateLimited(ctx context.Context, userID int64) bool {
	key := fmt.Sprintf("borrow-limit:%d", userID)
	uc.log.Info("key rate limiting: ", key)

	count, err := uc.redis.Client.Get(ctx, key).Int()
	if err != nil {
		uc.log.Error("error get rate limit: ", err)
	}

	uc.log.Info("total borrowing: ", count)
	if count >= 5 {
		return true
	}
	pipe := uc.redis.Client.TxPipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, 7*24*time.Hour)
	_, _ = pipe.Exec(ctx)
	return false
}

func mappingUpdateBook(dto *modelBook.Book, quantity int) (response dtoBook.UpdateBookRequest) {
	response = dtoBook.UpdateBookRequest{
		Quantity: quantity,
		Title:    dto.Title,
		ISBN:     dto.ISBN,
		Author:   dto.Author,
		Category: dto.Category,
	}

	return
}
