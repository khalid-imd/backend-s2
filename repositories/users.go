package repositories

import (
	"fundamental-golang/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
  CreateUser(user models.User) (models.User, error)
  FindUsers() ([]models.User, error)
  GetUser(ID int) (models.User, error)
  UpdateUser(user models.User, ID int) (models.User, error)
  DeleteUser(user models.User, ID int) (models.User, error)
}

type repository struct {
  db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
  return &repository{db} 
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
  err := r.db.Exec("INSERT INTO users(fullname,email,phone,location,image,role,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?)",user.Fullname,user.Email, user.Phone, user.Location, user.Image, user.Role, time.Now(), time.Now()).Error

  return user, err
}

func (r *repository) FindUsers() ([]models.User, error) {
  var users []models.User
  err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

  return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
  var user models.User
  err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

  return user, err
}

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
  err := r.db.Raw("UPDATE users SET fullname=?, email=?, phone=?, location=?, image=?, role=? WHERE id=?", user.Fullname,user.Email, user.Phone, user.Location, user.Image, user.Role, ID).Scan(&user).Error

  return user, err
}

func (r *repository) DeleteUser(user models.User,ID int) (models.User, error) {
  err := r.db.Raw("DELETE FROM users WHERE id=?",ID).Scan(&user).Error

  return user, err
}
