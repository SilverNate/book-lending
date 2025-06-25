package http

import (
	"book-lending-api/internal/borrow/dto"
	"book-lending-api/internal/borrow/usecase"
	"book-lending-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BorrowHandler struct {
	uc  usecase.IBorrowUseCase
	log *logrus.Logger
}

func NewBorrowHandler(uc usecase.IBorrowUseCase, log *logrus.Logger) *BorrowHandler {
	return &BorrowHandler{uc: uc, log: log}
}

func (h *BorrowHandler) Borrow(c *gin.Context) {
	userIDVal, ok := c.Get("userID")
	if !ok {
		response.Unauthorized(c, "Unauthorized")
		return
	}

	userID, ok := userIDVal.(int64)
	if !ok {
		response.Internal(c, "Invalid user id")
		return
	}

	var req dto.BorrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Warnf("request borrow books invalid: %v", err)
		response.BadRequest(c, "Invalid request")
		return
	}
	if err := h.uc.BorrowBooks(c.Request.Context(), userID, req); err != nil {
		h.log.Errorf("error handler borrow books: %v", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, "Book borrowed successfully")
}

func (h *BorrowHandler) Return(c *gin.Context) {
	var req dto.ReturnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Warnf("request return books invalid: %v", err)
		response.BadRequest(c, "Invalid request")
		return
	}
	if err := h.uc.ReturnBooks(c.Request.Context(), req); err != nil {
		h.log.Errorf("error handler return books: %v", err)
		response.Internal(c, err.Error())
		return
	}
	response.Success(c, "Book returned successfully")
}
