package dto

type CreateBookRequest struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	ISBN     string `json:"isbn" binding:"required"`
	Category string `json:"category"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
}

type UpdateBookRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}
