package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                  `json:"user_id" gorm:"type:varchar(25)"`
	User      UsersProfileResponse `json:"user"`
	Status    string               `json:"status"  gorm:"type:varchar(25)"`
	ProductID int                  `json:"product_id"`
	Product   ProductResponse      `json:"product"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}
