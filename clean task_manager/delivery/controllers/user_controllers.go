package controllers

import (
	"net/http"
	"task_manager/domain"

	"github.com/gin-gonic/gin"
)

func Register(UUC domain.UserUsecase)gin.HandlerFunc{
	return func (c *gin.Context){
		var user domain.UserInput

		err := c.ShouldBindJSON(&user)
		if err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return 
		}

		usr, er := UUC.Signup(user)
		if er != nil{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error":er.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"message":"User created Successfully.", "user" : usr})
	}
}

func Login(UUC domain.UserUsecase)gin.HandlerFunc{
	return func (c *gin.Context){
		var user domain.UserInput
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		usr, token, err := UUC.Login(user)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "user" : usr})
		}
	
}

func GetAllUsers(UUC domain.UserUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		users, err := UUC.GetUsers()
		if err != nil{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, users)
	}
}

func GetUserById(UUC domain.UserUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		user, err := UUC.GetUser(id)

		if err != nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"error" : err.Error()})
			return
		}

		c.IndentedJSON(http.StatusAccepted, user)
	}
}

func MakeAdmin(UUC domain.UserUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		user, err := UUC.MakeAdmin(id)

		if err != nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"error" : err.Error()})
			return
		}

		c.IndentedJSON(http.StatusAccepted, user)

	}
}


func DeleteUser(UUC domain.UserUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		err := UUC.DeleteUser(id)
		if err != nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"error" : err.Error()})
			return
		}

		c.IndentedJSON(http.StatusAccepted, gin.H{"message" : "User deleted successfully"})
	}
}


func UpdateUser(UUC domain.UserUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		var user domain.UserInput
		err := c.ShouldBindJSON(&user)
		if err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		usr, er := UUC.UpdateUser(id, user)
		if er != nil{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : er.Error()})
			return
		}
		c.IndentedJSON(http.StatusAccepted, usr)
	}
}