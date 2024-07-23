package service

import (
	"backend/dao"
	"backend/modals"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetSignup(ctx context.Context, userRequest modals.SignupRequest) error {

	var user modals.User

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return fmt.Errorf("error in decryting password %v", err)
	}
	user.Name = userRequest.Name
	user.Password = string(hash)
	user.Email = userRequest.Email
	user.Role = userRequest.Role
	err = dao.CreateUser(ctx, user)
	fmt.Println(user)
	//errr := dao.Login(ctx)
	if err != nil {
		return fmt.Errorf("Error in creating user")
	}
	return nil
}

func Login(ctx context.Context, loginRequest modals.LoginRequest) (*modals.User, error) {

	user, err := dao.Getuser(ctx, loginRequest.Email)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, fmt.Errorf("Wrong Password")
	}

	return user, nil
}
