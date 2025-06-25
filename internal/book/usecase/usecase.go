package usecase

import (
	"book-lending-api/internal/book/dto"
	"book-lending-api/internal/book/entity"
	"book-lending-api/internal/book/repository"
	log "book-lending-api/pkg/logger"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type BookUsecase struct {
	repo repository.IBookRepository
	log  *logrus.Logger
}

func NewBookUseCase(r repository.IBookRepository, log *logrus.Logger) *BookUsecase {
	return &BookUsecase{repo: r, log: log}
}

func (uc *BookUsecase) AddBook(ctx context.Context, input dto.CreateBookRequest) error {
	uc.log.WithField("addbook", input).Info("create book")

	existing, err := uc.repo.GetBookByTitleAndAuthor(ctx, input.Title, input.Author)
	if err != nil {
		return err
	}

	if existing != nil {
		uc.log.WithField("addbook", input).Info("book with same title and author already exists")
		return errors.New("book with same title and author already exists")
	}

	book := &entity.Book{
		Title:    input.Title,
		Author:   input.Author,
		ISBN:     input.ISBN,
		Category: input.Category,
		Quantity: input.Quantity,
	}
	return uc.repo.CreateBook(ctx, book)
}

func (uc *BookUsecase) GetAllBooks(ctx context.Context, offset, limit int) ([]entity.Book, error) {
	uc.log.WithField("getallbooks", fmt.Sprintf("%v, %v", offset, limit)).Info("get all book")
	return uc.repo.GetAllBook(ctx, offset, limit)
}

func (uc *BookUsecase) GetBookByID(ctx context.Context, id int64) (*entity.Book, error) {
	uc.log.WithField("getbookbyid", fmt.Sprintf("id: %v ", id)).Info("get book by id")
	return uc.repo.GetBookByID(ctx, id)
}

func (uc *BookUsecase) UpdateBook(ctx context.Context, id int64, input dto.UpdateBookRequest) error {
	uc.log.WithField("updatebook", fmt.Sprintf("request: %v", input)).Info("update book")

	book, err := uc.repo.GetBookByID(ctx, id)
	if err != nil {
		log.Error("book", "updatebook", fmt.Sprintf("error getting book, id: %v", id), err)
		return err
	}
	book.Title = input.Title
	book.Author = input.Author
	book.ISBN = input.ISBN
	book.Category = input.Category
	book.Quantity = input.Quantity
	return uc.repo.UpdateBook(ctx, book)
}

func (uc *BookUsecase) DeleteBook(ctx context.Context, id int64) error {
	uc.log.WithField("deletebook", fmt.Sprintf("id: %v ", id)).Info("delete book")
	return uc.repo.DeleteBook(ctx, id)
}
