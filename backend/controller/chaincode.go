package controller

import (
	//"backend/modals"
	//"backend/service"
	"backend/dlt"
	"backend/modals"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAllCoffee(c *gin.Context) {
	ctx := context.WithValue(c, ContextMap, c.Keys)
	result, err := dlt.GetAllCoffee(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Read all coffee batches successfully", "data": result})
}

func ReadCoffeeByBatchId(c *gin.Context) {
	ctx := context.WithValue(c, ContextMap, c.Keys)
	var createCoffee modals.DltResponse
	if err := c.ShouldBindJSON(&createCoffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err==>", err)
		return
	}
	result, err := dlt.GetCoffee(ctx, createCoffee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Read all coffee batches successfully", "data": result})
}

func CreateCoffee(c *gin.Context) {
	ctx := context.WithValue(c, ContextMap, c.Keys)
	var createCoffee modals.DltResponse
	if err := c.ShouldBindJSON(&createCoffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err :=  dlt.CreateCoffeeBatch(ctx, createCoffee)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": "coffee  batch is not created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create coffee  batch successfully"})
}
func UpdateCoffee(c *gin.Context) {
	ctx := context.WithValue(c, ContextMap, c.Keys)
	var updateCoffee modals.DltResponse
	if err := c.ShouldBindJSON(&updateCoffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err:=dlt.UpdateCoffeeBatch(ctx,updateCoffee)
	
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": "coffee  batch is not updated"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update coffee batch successfully"})
}

func TransferCoffee(c *gin.Context) {
	ctx := context.WithValue(c, ContextMap, c.Keys)
	var transferCoffee modals.DltResponse
	if err := c.ShouldBindJSON(&transferCoffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err:=dlt.TransferCoffeeOwnerShip(ctx,transferCoffee)
	
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": "coffee batch is not tranfered"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Trasnsfer coffee  batch ownership successfully"})
}
