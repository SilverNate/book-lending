package entity

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        int64 `gorm:"primaryKey"`
	Title     string
	Author    string
	ISBN      string
	Category  string
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
