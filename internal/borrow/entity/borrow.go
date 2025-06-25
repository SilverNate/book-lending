package entity

import (
	"time"
)

type Borrowing struct {
	ID         int64     `gorm:"primaryKey"`
	BookID     int64     `gorm:"not null"`
	UserID     int64     `gorm:"not null"`
	BorrowDate time.Time `gorm:"not null"`
	ReturnDate *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
