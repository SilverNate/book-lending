package repository

import (
	"book-lending-api/internal/book/entity"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

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
		r.log.Error("error create book: ", err)
		return err
	}
	return nil
}

func (r *BookRepository) GetAllBook(ctx context.Context, offset, limit int) ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Offset(offset).Limit(limit).Find(&books).Error
	if err != nil {
		r.log.Error("error get all books: ", err)
		return books, err
	}
	return books, err
}

func (r *BookRepository) GetBookByID(ctx context.Context, id int64) (*entity.Book, error) {
	var book entity.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		r.log.Error("error get book by id: ", err)
		return &book, err
	}
	return &book, err
}

func (r *BookRepository) UpdateBook(ctx context.Context, book *entity.Book) error {
	err := r.db.Save(book).Error
	if err != nil {
		r.log.Error("error update book: ", err)
	}
	return nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	err := r.db.Delete(&entity.Book{}, id).Error
	if err != nil {
		r.log.Error("error delete book: ", err)
		return err
	}
	return nil
}
