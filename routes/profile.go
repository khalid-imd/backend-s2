package routes

import (
	"fundamental-golang/handlers"
	"fundamental-golang/pkg/mysql"
	"fundamental-golang/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
  profileRepository := repositories.RepositoryProfile(mysql.DB)
  h := handlers.HandlerProfile(profileRepository)

  r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
}