package routes

import (
	"fundamental-golang/handlers"
	"fundamental-golang/pkg/mysql"
	"fundamental-golang/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
  productRepository := repositories.RepositoryProduct(mysql.DB)
  h := handlers.HandlerProduct(productRepository)

  r.HandleFunc("/product", h.CreateProduct).Methods("POST")
  r.HandleFunc("/products", h.FindProducts).Methods("GET")
  r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
  r.HandleFunc("/product/{id}", h.UpdateProduct).Methods("PATCH")
  r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
}