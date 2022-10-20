package main

import (
	"fmt"
	"fundamental-golang/database"
	"fundamental-golang/pkg/mysql"
	"fundamental-golang/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }
	
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()
	
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}