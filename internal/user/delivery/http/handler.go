package http

import (
	"book-lending-api/internal/user/dto"
	"book-lending-api/internal/user/usecase"
	"book-lending-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	uc  usecase.IUserUsecase
	log *logrus.Logger
}

func NewHandler(uc usecase.IUserUsecase, log *logrus.Logger) *Handler {
	return &Handler{uc: uc, log: log}
}

func (h *Handler) Register(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("register validation failed: ", err)
		response.BadRequest(c, "Invalid request body")
		return
	}
	err := h.uc.Register(c.Request.Context(), input)
	if err != nil {
		h.log.Error("registration failed: ", err)
		response.Internal(c, "Failed to register user")
		return
	}

	response.Created(c, "user registered successfully")
}

func (h *Handler) Login(c *gin.Context) {
	var input dto.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("login validation failed: ", err)
		response.BadRequest(c, "Invalid request body")
		return
	}
	token, err := h.uc.Login(c.Request.Context(), input)
	if err != nil {
		h.log.Error("Handler login failed: ", err)
		response.Unauthorized(c, "invalid credential")
		return
	}

	response.Success(c, token)
}
