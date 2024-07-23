package modals

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"  gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct{
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
