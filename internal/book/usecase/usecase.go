package usecase

import (
	"book-lending-api/internal/book/dto"
	"book-lending-api/internal/book/entity"
	"book-lending-api/internal/book/repository"
	"context"
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

func (uc *BookUsecase) Create(ctx context.Context, input dto.CreateBookRequest) error {
	uc.log.Infof("create book %+v", input)

	book := &entity.Book{
		Title:    input.Title,
		Author:   input.Author,
		ISBN:     input.ISBN,
		Category: input.Category,
		Quantity: input.Quantity,
	}
	return uc.repo.CreateBook(ctx, book)
}

func (uc *BookUsecase) GetAll(ctx context.Context, offset, limit int) ([]entity.Book, error) {
	uc.log.Info(fmt.Sprintf("get all book offset: %v, limit: %v", offset, limit))

	return uc.repo.GetAllBook(ctx, offset, limit)
}

func (uc *BookUsecase) GetByID(ctx context.Context, id int64) (*entity.Book, error) {
	uc.log.Infof("get book %+v", id)
	return uc.repo.GetBookByID(ctx, id)
}

func (uc *BookUsecase) Update(ctx context.Context, id int64, input dto.UpdateBookRequest) error {
	uc.log.Infof("update book %+v", input)

	book, err := uc.repo.GetBookByID(ctx, id)
	if err != nil {
		uc.log.Error("error getting book by id: ", err)
		return err
	}
	book.Title = input.Title
	book.Author = input.Author
	book.ISBN = input.ISBN
	book.Category = input.Category
	book.Quantity = input.Quantity
	return uc.repo.UpdateBook(ctx, book)
}

func (uc *BookUsecase) Delete(ctx context.Context, id int64) error {
	uc.log.Infof("delete book %+v", id)
	return uc.repo.DeleteBook(ctx, id)
}
