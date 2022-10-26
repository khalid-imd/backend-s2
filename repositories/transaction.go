package repositories

import (
	"fundamental-golang/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(user models.Transaction) (models.Transaction, error)
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	DeleteTransaction(user models.Transaction, ID int) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	// err := r.db.Exec("INSERT INTO transactions(product_id,product,user_id,user,price,status,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?)",transaction.ProductID,transaction.Product, transaction.UserID, transaction.User, transaction.Price, transaction.Status, time.Now(), time.Now()).Error
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Product").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, ID).Error // add this code

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Raw("DELETE FROM transactions WHERE id=?", ID).Scan(&transaction).Error

	return transaction, err
}
