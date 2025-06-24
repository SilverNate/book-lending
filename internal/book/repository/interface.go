package repository

import (
	"book-lending-api/internal/book/entity"
	"context"
)

type IBookRepository interface {
	CreateBook(ctx context.Context, book *entity.Book) error
	GetAllBook(ctx context.Context, offset, limit int) ([]entity.Book, error)
	GetBookByID(ctx context.Context, id int64) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) error
	DeleteBook(ctx context.Context, id int64) error
}
