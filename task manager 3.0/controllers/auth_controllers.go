package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func Register(tm *data.Taskmanager)gin.HandlerFunc{
	return func (c *gin.Context){
		var user models.User

		err := c.ShouldBindJSON(&user)
		if err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
			return 
		}

		er := tm.Signup(user)
		if er != nil{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message":er.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, user)
	}
}

func Enter(tm *data.Taskmanager)gin.HandlerFunc{
	return func (c *gin.Context){
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		token, err := tm.Login(user)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Internal server error"})
		}

		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
		}
	
}

