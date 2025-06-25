package db

import (
	"book-lending-api/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL(cfg *config.EnvConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DBDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
