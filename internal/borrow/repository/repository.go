package repository

import (
	"book-lending-api/internal/borrow/entity"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type BorrowRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewBorrowRepository(db *gorm.DB, log *logrus.Logger) *BorrowRepository {
	return &BorrowRepository{db: db, log: log}
}

func (r *BorrowRepository) CreateBorrowing(ctx context.Context, b *entity.Borrowing) error {
	err := r.db.WithContext(ctx).Create(b).Error
	if err != nil {
		r.log.Error("error create borrowing: ", err)
		return err
	}
	return nil
}

func (r *BorrowRepository) FindBorrowingByID(ctx context.Context, id int64) (*entity.Borrowing, error) {
	var borrow entity.Borrowing
	err := r.db.WithContext(ctx).First(&borrow, id).Error
	if err != nil {
		r.log.Error("error find borrowing : , id: ", err, id)
		return &borrow, err
	}
	return &borrow, err
}

func (r *BorrowRepository) FindActiveByUserAndBook(ctx context.Context, userID, bookID int64) (*entity.Borrowing, error) {
	var borrow entity.Borrowing
	err := r.db.WithContext(ctx).Where("user_id = ? AND book_id = ? AND return_date IS NULL", userID, bookID).First(&borrow).Error
	if err != nil {
		r.log.Error("error find borrowing by active user and book: ", err, userID, bookID)
		return &borrow, err
	}
	return &borrow, err
}

func (r *BorrowRepository) InsertOrUpdateBorrowing(ctx context.Context, borrow *entity.Borrowing) error {
	err := r.db.WithContext(ctx).Save(borrow).Error
	if err != nil {
		r.log.Error("error insert or update borrowing: ", err)
		return err
	}
	return nil
}

func (r *BorrowRepository) CountBorrowsInLast7Days(ctx context.Context, userID int64) (int, error) {
	var count int64
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	err := r.db.WithContext(ctx).
		Model(&entity.Borrowing{}).
		Where("user_id = ? AND borrow_date >= ?", userID, sevenDaysAgo).
		Count(&count).Error
	if err != nil {
		r.log.Error("error count borrowings: ", err)
		return 0, err
	}
	return int(count), err
}
