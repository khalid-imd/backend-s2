package authdto

type RegisterResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"name"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Role     string `gorm:"type: varchar(255)" json:"role"`
}
type LoginResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
