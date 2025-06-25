package http

import (
	"book-lending-api/internal/book/dto"
	"book-lending-api/internal/book/repository"
	"book-lending-api/internal/book/usecase"
	"book-lending-api/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Handler struct {
	uc  usecase.IBookUseCase
	log *logrus.Logger
}

func NewHandler(uc usecase.IBookUseCase, log *logrus.Logger) *Handler {
	return &Handler{uc: uc, log: log}
}

// CreateBooks godoc
// @Summary Create a new book
// @Description Authenticated user borrows a book
// @Tags Book
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param book body dto.CreateBookRequest true "Book Data"
// @Success 200 {object} response.APIResponse
// @Router /books [post]
func (h *Handler) CreateBooks(c *gin.Context) {
	var input dto.CreateBookRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("create book request is invalid: ", err)
		response.BadRequest(c, "Invalid request")
		return
	}
	if err := h.uc.AddBook(c.Request.Context(), input); err != nil {
		h.log.Error("error creating book: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Created(c, "book created successfully")
}

// GetAllBooks godoc
// @Summary List all books
// @Description Authenticated user borrows a book
// @Tags Book
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} response.APIResponse
// @Router /books [get]
func (h *Handler) GetAllBooks(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	books, err := h.uc.GetAllBooks(c.Request.Context(), offset, limit)
	if err != nil {
		h.log.Error("error getting all books: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, books)
}

// GetBooksByID godoc
// @Summary Get book by id
// @Description Authenticated user borrows a book
// @Tags Book
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.APIResponse
// @Router /books/id [get]
func (h *Handler) GetBooksByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	book, err := h.uc.GetBookByID(c.Request.Context(), id)
	if err != nil {
		h.log.Error("error getting book by id: ", err)
		switch {
		case errors.Is(err, repository.ErrBookNotFound):
			response.NotFound(c, "book not found")
		default:
			response.Internal(c, "internal server error")
		}

		response.BadRequest(c, "Book not found")
		return
	}
	response.Success(c, book)
}

// UpdateBooks godoc
// @Summary Update book by id
// @Description Authenticated user borrows a book
// @Tags Book
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.APIResponse
// @Router /books/id [put]
func (h *Handler) UpdateBooks(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var input dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("update book request is invalid: ", err)
		response.BadRequest(c, "Invalid request")
		return
	}
	if err := h.uc.UpdateBook(c.Request.Context(), id, input); err != nil {
		h.log.Error("error updating book: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, "book updated successfully")
}

// DeleteBooks godoc
// @Summary Delete book by id
// @Description Authenticated user borrows a book
// @Tags Book
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.APIResponse
// @Router /books/id [delete]
func (h *Handler) DeleteBooks(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.uc.DeleteBook(c.Request.Context(), id); err != nil {
		h.log.Error("error deleting book: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, "book deleted successfully")
}
