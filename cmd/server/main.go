package main

import (
	"book-lending-api/config"
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

	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
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
