package controllers

import (
	"context"
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(tm *data.Taskmanager) gin.HandlerFunc{
		return func (c *gin.Context){
			tasks, err := tm.GetTasks(context.TODO())
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
		
		task, err := tm.GetTask(context.TODO(),id)
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
			err := tm.PostTask(context.TODO(),task)
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
		err := tm.DeleteTask(context.TODO(),id)
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
		task, err := tm.UpdateTask(context.TODO(), id, task)
		if err == nil{
			c.IndentedJSON(http.StatusOK, task)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "task not found"})
	}
}