package http

import (
	"book-lending-api/internal/book/dto"
	"book-lending-api/internal/book/usecase"
	"book-lending-api/pkg/response"
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

func (h *Handler) CreateBooks(c *gin.Context) {
	var input dto.CreateBookRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("create book request is invalid: ", err)
		response.BadRequest(c, "Invalid request")
		return
	}
	if err := h.uc.Create(c.Request.Context(), input); err != nil {
		h.log.Error("error creating book: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Created(c, "book created successfully")
}

func (h *Handler) GetAllBooks(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	books, err := h.uc.GetAll(c.Request.Context(), offset, limit)
	if err != nil {
		h.log.Error("error getting all books: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, books)
}

func (h *Handler) GetBooksByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	book, err := h.uc.GetByID(c.Request.Context(), id)
	if err != nil {
		h.log.Error("error getting book by id: ", err)
		response.BadRequest(c, "Book not found")
		return
	}
	response.Success(c, book)
}

func (h *Handler) UpdateBooks(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var input dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("update book request is invalid: ", err)
		response.BadRequest(c, "Invalid request")
		return
	}
	if err := h.uc.Update(c.Request.Context(), id, input); err != nil {
		h.log.Error("error updating book: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, "book updated successfully")
}

func (h *Handler) DeleteBooks(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.uc.Delete(c.Request.Context(), id); err != nil {
		h.log.Error("error deleting book: ", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, "book deleted successfully")
}
