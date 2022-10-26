package transactiondto

type CreateTransactionRequest struct {
	UserID    int    `json:"user_id" form:"user_id"`
	Status    string `json:"status" form:"status" validate:"required"`
	ProductID int    `json:"product_id" form:"product_id" validate:"required"`
}

type UpdateTransactionRequest struct {
	ProductID int    `json:"product_id" form:"product_id" validate:"required"`
	UserID    int    `json:"user_id" form:"user_id"`
	Price     int    `json:"price" form:"price"`
	Status    string `json:"status" form:"status"`
}
