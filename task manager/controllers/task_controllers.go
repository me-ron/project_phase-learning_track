package controllers

import (
	"net/http"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context, tasks []*models.Task){
	c.IndentedJSON(http.StatusOK, tasks)

}

func GetTaskById(c *gin.Context, tasks []*models.Task){
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

func PostTask(c *gin.Context, tasks []*models.Task){
	var task models.Task

	if err := c.BindJSON(&task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
	}

	tasks = append(tasks, &task)
	c.IndentedJSON(http.StatusCreated, *tasks[len(tasks) - 1])
}

func DeleteTask(c *gin.Context, tasks []*models.Task){
	id := c.Param("id")
	for i := range len(tasks){
		if tasks[i].ID == id{
			tasks = append(tasks[:i], tasks[i + 1:]...)
			c.IndentedJSON(http.StatusAccepted, gin.H{"messages" : "deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "no task of this id"})
}

func UpdateTask(c *gin.Context, tasks []*models.Task){
	var task models.Task

	if err := c.BindJSON(&task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
	} 

	for i := range len(tasks){
		if tasks[i].ID == task.ID{
			tasks[i] = &task
			c.IndentedJSON(http.StatusAccepted, task)
		}
	}
}
