package repository

import (
	"book-lending-api/internal/book/entity"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

var ErrBookNotFound = errors.New("book not found")

type BookRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewBookRepository(db *gorm.DB, log *logrus.Logger) *BookRepository {
	return &BookRepository{db: db, log: log}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *entity.Book) error {
	err := r.db.Create(book).Error
	if err != nil {
		r.log.WithField("createbook", fmt.Sprintf("request: %v ", book)).Error("create book")
		return err
	}
	return nil
}

func (r *BookRepository) GetAllBook(ctx context.Context, offset, limit int) ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Offset(offset).Limit(limit).Find(&books).Error
	if err != nil {
		r.log.WithField("getallbook", fmt.Sprintf("offset & limit: %v, %v", offset, limit)).Error("get all book")
		return books, err
	}
	return books, err
}

func (r *BookRepository) GetBookByID(ctx context.Context, id int64) (*entity.Book, error) {
	var book entity.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		r.log.WithField("getbookby-id", fmt.Sprintf("id: %v ", id)).Error("get book by id")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookNotFound
		}
		return &book, fmt.Errorf("get book by id failed: %w", err)
	}
	return &book, err
}

func (r *BookRepository) UpdateBook(ctx context.Context, book *entity.Book) error {
	err := r.db.Save(book).Error
	if err != nil {
		r.log.WithField("updatebook", fmt.Sprintf("request: %v ", book)).Error("update book")
	}
	return nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	err := r.db.Delete(&entity.Book{}, id).Error
	if err != nil {
		r.log.WithField("deletebook", fmt.Sprintf("id: %v", id)).Error("delete book")
		return err
	}
	return nil
}

func (r *BookRepository) GetBookByTitleAndAuthor(ctx context.Context, title, author string) (*entity.Book, error) {
	var book entity.Book

	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)

	err := r.db.WithContext(ctx).
		Where("LOWER(title) = ? AND LOWER(author) = ?", strings.ToLower(title), strings.ToLower(author)).
		First(&book).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &book, nil
}
