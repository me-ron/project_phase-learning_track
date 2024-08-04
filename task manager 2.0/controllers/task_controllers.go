package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(tm *data.Taskmanager) gin.HandlerFunc{
		return func (c *gin.Context){
			tasks, err := tm.Get_tasks()
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
		
		task, err := tm.Get_task(id)
		if err != nil{
			c.IndentedJSON(http.StatusOK, task)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Task not found"})

	}
}

func PostTask(tm *data.Taskmanager) gin.HandlerFunc{
	return func (c *gin.Context){
			var task models.Task

			if err := c.BindJSON(&task); err != nil{
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
				return
			}
			err := tm.Post_Task(task)
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
		err := tm.Delete_task(id)
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
		task, err := tm.Update_task(id, task)
		if err == nil{
			c.IndentedJSON(http.StatusOK, task)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "task not found"})
	}
}