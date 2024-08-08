package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func Register(tm *data.Taskmanager)gin.HandlerFunc{
	return func (c *gin.Context){
		var user models.UserInput

		err := c.ShouldBindJSON(&user)
		if err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
			return 
		}

		usr, er := tm.Signup(user)
		if er != nil{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message":er.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"msg":"User created Successfully.", "user" : models.ChangeToOutput(usr)})
	}
}

func Login(tm *data.Taskmanager)gin.HandlerFunc{
	return func (c *gin.Context){
		var user models.UserInput
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		token, err := tm.Login(user)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "user" : models.ChangeToOutput(user)})
		}
	
}

