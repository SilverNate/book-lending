package repository

import (
	"book-lending-api/internal/borrow/entity"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BorrowRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewBorrowRepository(db *gorm.DB, log *logrus.Logger) *BorrowRepository {
	return &BorrowRepository{db: db, log: log}
}

func (r *BorrowRepository) CreateBorrowing(ctx context.Context, borrow *entity.Borrowing) error {
	err := r.db.WithContext(ctx).Create(borrow).Error
	if err != nil {
		r.log.WithField("createborrowing", fmt.Sprintf("request: %v", borrow)).Error("create borrowing")
		return err
	}
	return nil
}

func (r *BorrowRepository) GetBorrowingByID(ctx context.Context, id int64) (*entity.Borrowing, error) {
	var borrow entity.Borrowing
	err := r.db.WithContext(ctx).First(&borrow, id).Error
	if err != nil {
		r.log.WithField("getborrowingbyid", fmt.Sprintf("id:  %v", id)).Error("get borrowing by id")
		return &borrow, err
	}
	return &borrow, err
}

func (r *BorrowRepository) IsBookBorrowed(ctx context.Context, userID, bookID int64) (bool, error) {
	var borrow entity.Borrowing

	err := r.db.WithContext(ctx).
		Where("user_id = ? AND book_id = ? AND return_date IS NULL", userID, bookID).
		First(&borrow).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithFields(logrus.Fields{
			"book_id": bookID,
			"user_id": userID,
		}).Info("no active borrow found")
		return false, nil
	}

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"book_id": bookID,
			"user_id": userID,
			"error":   err,
		}).Error("error checking active borrow")
		return false, err
	}

	return true, nil
}

func (r *BorrowRepository) InsertOrUpdateBorrowing(ctx context.Context, borrow *entity.Borrowing) error {
	err := r.db.WithContext(ctx).Save(borrow).Error
	if err != nil {
		r.log.WithField("insertorupdateborrowing", fmt.Sprintf("request:  %v", borrow)).Error("insert or update borrowing")
		return err
	}
	return nil
}
