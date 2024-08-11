package controllers

import (
	"net/http"
	"task_manager/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasks(TUC domain.TaskUsecase) gin.HandlerFunc{
		return func (c *gin.Context){
			f, exists := c.Get("filter")
			if !exists{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "filter couldn't be found"})
				return
			}

			filter, ok := f.(bson.M)
			if !ok{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "type assertion didn't work"})
				return
			}
			
			tasks, err := TUC.GetTasks(filter)
			if err != nil{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : err.Error()})
				return
			}
			c.IndentedJSON(http.StatusOK, tasks)

		}
}

func GetTaskById(TUC domain.TaskUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		user, exists := c.Get("user")
		if !exists{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "user couldn't be found"})
			return
		}

		usr, ok := user.(domain.DBUser)
		if !ok{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "type assertion didn't work"})
			return
		}
		
		task, err := TUC.GetTask(id, usr.ID)
		if err == nil{
			c.IndentedJSON(http.StatusOK, task)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"error" : err.Error()})
	}
}

func PostTask(TUC domain.TaskUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
			var task domain.Task

			if err := c.BindJSON(&task); err != nil{
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request"})
				return
			}
			user, exists := c.Get("user")
			if !exists{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "user couldn't be found"})
				return
			}

			usr, ok := user.(domain.DBUser)
			if !ok{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "type assertion didn't work"})
				return
			}

			task, err := TUC.PostTask(task, usr)
			if err != nil{
				c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : err.Error()})
				return
			}
			c.IndentedJSON(http.StatusCreated, gin.H{"message" : "created sucessfully", "task": task})
		}

}


func DeleteTask(TUC domain.TaskUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		user, exists := c.Get("user")
		if !exists{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "user couldn't be found"})
			return
		}

		usr, ok := user.(domain.DBUser)
		if !ok{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "type assertion didn't work"})
			return
		}
		err := TUC.DeleteTask(id, usr.ID)
		if err == nil{
			c.IndentedJSON(http.StatusOK, gin.H{"messages" : "deleted successfully"})
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"error" : err.Error()})
	}

}

func UpdateTask(TUC domain.TaskUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		var task domain.Task

		if err := c.BindJSON(&task); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request"})
		} 
		
		id := c.Param("id")
		user, exists := c.Get("user")
		if !exists{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "user couldn't be found"})
			return
		}

		usr, ok := user.(domain.DBUser)
		if !ok{
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error" : "type assertion didn't work"})
			return
		}

		task, err := TUC.UpdateTask(id, task, usr)
		if err == nil{
			c.IndentedJSON(http.StatusOK, task)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"error" : "task not found"})
	}
}