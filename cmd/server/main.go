package main

import (
	"book-lending-api/config"
	"book-lending-api/internal/book"
	"book-lending-api/internal/borrow"
	"book-lending-api/internal/middleware"
	"book-lending-api/internal/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.LoadEnv()

	userHandler := user.InitUserHandler()

	r := gin.Default()
	jwtService := middleware.NewJWTService(cfg)
	authMiddleware := middleware.JWTMiddleware(jwtService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	bookHandler := book.InitBookHandler()
	books := r.Group("/books", authMiddleware)
	{
		books.POST("", bookHandler.CreateBooks)
		books.GET("", bookHandler.GetAllBooks)
		books.GET(":id", bookHandler.GetBooksByID)
		books.PUT(":id", bookHandler.UpdateBooks)
		books.DELETE(":id", bookHandler.DeleteBooks)
	}

	borrowHandler := borrow.InitBorrowHandler()
	borrowPath := r.Group("/borrowing", authMiddleware)
	{
		borrowPath.POST("/borrow", borrowHandler.Borrow)
		borrowPath.POST("/return", borrowHandler.Return)
	}

	s := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server running on port", cfg.Port)
	log.Fatal(s.ListenAndServe())
}
