package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasks(tm *data.Taskmanager) gin.HandlerFunc{
		return func (c *gin.Context){
			f, exists := c.Get("filter")
			if !exists{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "filter couldn't be found"})
				return
			}

			filter, ok := f.(bson.M)
			if !ok{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "type assertion didn't work"})
				return
			}
			
			tasks, err := tm.GetTasks(filter)
			if err != nil{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : err.Error()})
				return
			}
			c.IndentedJSON(http.StatusOK, tasks)

		}
}

func GetTaskById(tm *data.Taskmanager) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		user, exists := c.Get("user")
		if !exists{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "user couldn't be found"})
			return
		}

		usr, ok := user.(models.DBUser)
		if !ok{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "type assertion didn't work"})
			return
		}
		
		task, err := tm.GetTask(id, usr.ID)
		if err == nil{
			c.IndentedJSON(http.StatusOK, task)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : err.Error()})

	}
}

func PostTask(tm *data.Taskmanager) gin.HandlerFunc{
	return func (c *gin.Context){
			var task models.Task

			if err := c.BindJSON(&task); err != nil{
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
				return
			}
			user, exists := c.Get("user")
			if !exists{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "user couldn't be found"})
				return
			}

			usr, ok := user.(models.DBUser)
			if !ok{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "type assertion didn't work"})
				return
			}
			task.User = usr
			err := tm.PostTask(task)
			if err != nil{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : err.Error()})
				return
			}
			c.IndentedJSON(http.StatusCreated, gin.H{"message" : "created sucessfully"})
		}

}


func DeleteTask(tm *data.Taskmanager) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		user, exists := c.Get("user")
		if !exists{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "user couldn't be found"})
			return
		}

		usr, ok := user.(models.DBUser)
		if !ok{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "type assertion didn't work"})
			return
		}
		err := tm.DeleteTask(id, usr.ID)
		if err == nil{
			c.IndentedJSON(http.StatusOK, gin.H{"messages" : "deleted successfully"})
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "no task of this id"})
	}

}

func UpdateTask(tm *data.Taskmanager) gin.HandlerFunc{
	return func (c *gin.Context){
		var task models.Task

		if err := c.BindJSON(&task); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
		} 
		
		id := c.Param("id")
		user, exists := c.Get("user")
		if !exists{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "user couldn't be found"})
			return
		}

		usr, ok := user.(models.DBUser)
		if !ok{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message" : "type assertion didn't work"})
			return
		}

		task, err := tm.UpdateTask(id, task, usr)
		if err == nil{
			c.IndentedJSON(http.StatusOK, task)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "task not found"})
	}
}