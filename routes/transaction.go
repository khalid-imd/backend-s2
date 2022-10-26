package routes

import (
	"fundamental-golang/handlers"
	"fundamental-golang/pkg/middleware"
	"fundamental-golang/pkg/mysql"
	"fundamental-golang/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.UpdateTransaction).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
}
