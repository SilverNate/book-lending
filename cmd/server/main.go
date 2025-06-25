package main

import (
	"book-lending-api/config"
	"book-lending-api/internal/book"
	"book-lending-api/internal/borrow"
	"book-lending-api/internal/middleware"
	"book-lending-api/internal/user"
	"book-lending-api/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.LoadEnv()

	userHandler := user.InitUserHandler()
	infrastructure.SetupElasticLogger()
	_, err := infrastructure.InitInfra()
	if err != nil {
		log.Fatalf("failed to initialize redis: %v", err)
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		borrowPath.POST("/borrow", borrowHandler.BorrowBook)
		borrowPath.POST("/return", borrowHandler.ReturnBook)
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
