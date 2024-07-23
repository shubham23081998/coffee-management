package dao

import (
	"backend/config"
	"backend/modals"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type userDao struct {
	db *gorm.DB
}
type DBWrapper struct {
	db *gorm.DB
}

var body struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
var db *gorm.DB

func NewUserDao(da *DBWrapper) *userDao {
	return &userDao{db: da.db}
}

func CreateUser(ctx context.Context, users modals.User) error {
	
	user := &modals.User{Name: users.Name, Email: users.Email, Password: users.Password,Role: users.Role}
	result := config.DB.Create(&user)
	fmt.Println(result)
	if result.Error != nil {
		return fmt.Errorf("error in creating user %v", result.Error)
	}
	return nil
}

func Getuser(ctx context.Context, email string) (*modals.User,error) {
	var user modals.User
	config.DB.First(&user, "email= ?", email)
	if user.ID == 0 {
		return &modals.User{},fmt.Errorf("Invalid email or password")

	}
	return &user,nil
}
