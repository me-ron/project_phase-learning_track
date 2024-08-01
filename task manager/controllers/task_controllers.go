package controllers

import (
	"net/http"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(tasks []*models.Task) gin.HandlerFunc{
		return func (c *gin.Context){
			c.IndentedJSON(http.StatusOK, tasks)

		}
}

func GetTaskById(tasks []*models.Task) gin.HandlerFunc{

	return func (c *gin.Context){
		id := c.Param("id")
		for i := range len(tasks){
			if tasks[i].ID == id {
				task := *tasks[i]
				c.IndentedJSON(http.StatusOK, task)
				return
			}
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Taks not found"})

	}
}

func PostTask(tasks []*models.Task) gin.HandlerFunc{
	return func (c *gin.Context){
		var task models.Task

		if err := c.BindJSON(&task); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
		}

		tasks = append(tasks, &task)
		c.IndentedJSON(http.StatusCreated, *tasks[len(tasks) - 1])
	}

}


func DeleteTask(tasks []*models.Task) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		for i := range len(tasks){
			if tasks[i].ID == id{
				tasks = append(tasks[:i], tasks[i + 1:]...)
				c.IndentedJSON(http.StatusOK, gin.H{"messages" : "deleted successfully"})
				return
			}
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "no task of this id"})
	}

}

func UpdateTask(tasks []*models.Task) gin.HandlerFunc{
	return func (c *gin.Context){
		var task models.Task

		if err := c.BindJSON(&task); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
		} 

		for i := range len(tasks){
			if tasks[i].ID == task.ID{
				tasks[i] = &task
				c.IndentedJSON(http.StatusOK, task)
			}
		}
	}
}