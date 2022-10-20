package database

import (
	"fmt"
	"fundamental-golang/models"
	"fundamental-golang/pkg/mysql"
)

func RunMigration() {
  err := mysql.DB.AutoMigrate(&models.User{},&models.Product{})

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}