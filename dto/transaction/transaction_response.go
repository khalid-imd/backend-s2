package transactiondto

type TransactionResponse struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id" form:"product_id" validate:"required"`
	Product   string `json:"product" form:"product" validate:"required"`
	UserID    int    `json:"user_id" form:"user_id" validate:"required"`
	User      string `json:"user" form:"user" validate:"required"`
	Price     int    `json:"price" form:"price" validate:"required"`
	Status    string `json:"status" form:"status" validate:"required"`
}