package routes

import (
	// "backend/controller"

	"backend/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	// router.GET("user/", controller.GetUsers)
	// router.POST("user/", controller.CreateUser)
	// router.DELETE("user/:id", controller.DeleteUser)
	// router.PUT("user/:id", controller.UpdateUser)
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.GET("/getAllCoffee", controller.ReadAllCoffee)
	router.GET("/getCoffeeByBatchId", controller.ReadCoffeeByBatchId)
	router.POST("/createCoffee", controller.CreateCoffee)
	router.POST("/updatecoffee", controller.UpdateCoffee)
	router.POST("/trasnsfercoffee", controller.TransferCoffee)

}
