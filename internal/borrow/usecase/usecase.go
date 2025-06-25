package usecase

import (
	"book-lending-api/internal/book/usecase"
	"book-lending-api/internal/borrow/dto"
	"book-lending-api/internal/borrow/entity"
	"book-lending-api/internal/borrow/repository"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

type BorrowUsecase struct {
	repo        repository.IBorrowRepository
	log         *logrus.Logger
	bookUsecase usecase.IBookUseCase
}

func NewBorrowUsecase(repo repository.IBorrowRepository, log *logrus.Logger, bookUsecase usecase.IBookUseCase) *BorrowUsecase {
	return &BorrowUsecase{repo: repo, log: log, bookUsecase: bookUsecase}
}

func (uc *BorrowUsecase) BorrowBooks(ctx context.Context, userID int64, req dto.BorrowRequest) error {
	uc.log.Infof("borrowing books userId: %v, request: %v", userID, req)

	book, err := uc.bookUsecase.GetBookByID(ctx, req.BookID)
	if err != nil || book == nil {
		uc.log.Errorf("error get book when borrowing: %v", err)
		return errors.New("book not found")
	}

	count, err := uc.repo.CountBorrowsInLast7Days(ctx, userID)
	if err != nil {
		return err
	}
	if count >= 5 {
		return errors.New("borrow limit exceeded (max 5 per 7 days)")
	}
	_, err = uc.repo.FindActiveByUserAndBook(ctx, userID, req.BookID)
	if err == nil {
		return errors.New("book already borrowed and not returned")
	}

	borrow := &entity.Borrowing{
		BookID:     req.BookID,
		UserID:     userID,
		BorrowDate: time.Now(),
	}
	return uc.repo.CreateBorrowing(ctx, borrow)
}

func (uc *BorrowUsecase) ReturnBooks(ctx context.Context, req dto.ReturnRequest) error {
	uc.log.Infof("return books reqquest: %v", req)

	borrow, err := uc.repo.FindBorrowingByID(ctx, req.BorrowingID)
	if err != nil {
		return errors.New("borrowing record not found")
	}
	now := time.Now()
	borrow.ReturnDate = &now
	return uc.repo.InsertOrUpdateBorrowing(ctx, borrow)
}
