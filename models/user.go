package models

import "time"

type User struct {
  ID        int                   `json:"id"`
  Fullname  string                `json:"fullname" gorm:"type: varchar(255)"`
  Email     string                `json:"email" gorm:"type: varchar(255)"`
  Phone     string                `json:"phone" gorm:"type: varchar(255)"` 
  Location  string                `json:"location" gorm:"type: varchar(255)"` 
  Image     string                `json:"image" gorm:"type: varchar(255)"`
  Role      string                `json:"role" gorm:"type: varchar(255)"`
//   Profile   ProfileResponse       `json:"profile"`
//   Products  []ProductUserResponse `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  CreatedAt time.Time             `json:"-"`
  UpdatedAt time.Time             `json:"-"`
}

type UsersProfileResponse struct {
  ID   int    `json:"id"`
  Fullname string `json:"fullname"`
  Email string `json:"email"`
}

func (UsersProfileResponse) TableName() string {
  return "users"
}