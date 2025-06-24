package usecase_test

import (
	"book-lending-api/internal/book/dto"
	"book-lending-api/internal/book/entity"
	"book-lending-api/internal/book/mocks"
	"book-lending-api/internal/book/usecase"
	"context"
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBook_Success(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	log := logrus.New()
	bookUC := usecase.NewBookUseCase(mockRepo, log)

	req := dto.CreateBookRequest{
		Title:    "Domain-Driven Design",
		Author:   "Eric Evans",
		ISBN:     "9780321125217",
		Category: "Design",
		Quantity: 3,
	}

	mockRepo.On("CreateBook", mock.Anything, mock.Anything).Return(nil)

	err := bookUC.Create(context.Background(), req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllBook_Success(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	log := logrus.New()
	bookUC := usecase.NewBookUseCase(mockRepo, log)

	mockRepo.On("GetAllBook", mock.Anything, 0, 10).Return([]entity.Book{
		{ID: 1, Title: "Book 1"},
	}, nil)

	books, err := bookUC.GetAll(context.Background(), 0, 10)

	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Book 1", books[0].Title)
}

func TestUpdateBook_Success(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	log := logrus.New()
	bookUC := usecase.NewBookUseCase(mockRepo, log)

	book := &entity.Book{
		ID:       1,
		Title:    "Old Title",
		Quantity: 1,
	}

	mockRepo.On("GetBookByID", mock.Anything, int64(1)).Return(book, nil)
	mockRepo.On("UpdateBook", mock.Anything, mock.Anything).Return(nil)

	err := bookUC.Update(context.Background(), 1, dto.UpdateBookRequest{
		Title:    "New Title",
		Author:   "New Author",
		ISBN:     "123456",
		Category: "New Cat",
		Quantity: 5,
	})

	assert.NoError(t, err)
	assert.Equal(t, "New Title", book.Title)
	mockRepo.AssertExpectations(t)
}
