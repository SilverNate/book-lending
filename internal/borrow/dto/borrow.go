package dto

type BorrowRequest struct {
	BookID int64 `json:"book_id" binding:"required"`
}

type ReturnRequest struct {
	BorrowingID int64 `json:"borrowing_id" binding:"required"`
}
