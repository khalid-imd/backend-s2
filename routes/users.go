package routes

import (
	"fundamental-golang/handlers"
	"fundamental-golang/pkg/mysql"
	"fundamental-golang/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
  userRepository := repositories.RepositoryUser(mysql.DB)
  h := handlers.HandlerUser(userRepository)

  r.HandleFunc("/user", h.CreateUser).Methods("POST")
  r.HandleFunc("/user", h.FindUsers).Methods("GET")
  r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
  r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
  r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
