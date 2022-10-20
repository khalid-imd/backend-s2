package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
  TodoRoutes(r)
  UserRoutes(r)
  ProfileRoutes(r)
  ProductRoutes(r)
  AuthRoutes(r)
}