package controller

import (
	"backend/modals"
	"backend/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ContextMap string = "context_map"

func Signup(c *gin.Context) {
	ctx := context.WithValue(c, ContextMap, c.Keys)
	var userInput modals.SignupRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.GetSignup(ctx, userInput)

	c.JSON(http.StatusOK, gin.H{"message": "signup successfully"})

}

func Login(c *gin.Context) {
	ctx := context.WithValue(c, ContextMap, c.Keys)
	var loginReuest modals.LoginRequest
	if err := c.ShouldBindJSON(&loginReuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := service.Login(ctx, loginReuest)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully", "data": user})
}
