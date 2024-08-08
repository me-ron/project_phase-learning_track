package controllers

// import (
// 	"net/http"
// 	"task_manager/data"
// 	"task_manager/models"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func GetAllUsers(tm *data.Taskmanager) gin.HandlerFunc{
// 	return func (c *gin.Context){
// 		users, err := tm.GetUsers()
// 		if err != nil{
// 			c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : err.Error()})
// 			return
// 		}

// 		c.IndentedJSON(http.StatusOK, users)
// 	}
// }

// func GetUserById(tm *data.Taskmanager) gin.HandlerFunc{
// 	return func (c *gin.Context){
// 		id := c.Param("id")
// 		user, err := tm.GetUser(id)

// 		if err != nil{
// 			c.IndentedJSON(http.StatusNotFound, gin.H{"message" : err.Error()})
// 			return
// 		}

// 		c.IndentedJSON(http.StatusAccepted, user)
// 	}
// }

// func MakeAdmin(tm *data.Taskmanager) gin.HandlerFunc{
// 	return func (c *gin.Context){
// 		id := c.Param("id")
// 		user, err := tm.ChangeRole(id)

// 		if err != nil{
// 			c.IndentedJSON(http.StatusNotFound, gin.H{"message" : err.Error()})
// 			return
// 		}

// 		c.IndentedJSON(http.StatusAccepted, user)

// 	}
// }

