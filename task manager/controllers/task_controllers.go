package controllers

import (
	"log"
	"net/http"
	"strconv"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(tasks *[]*models.Task) gin.HandlerFunc{
		return func (c *gin.Context){
			c.IndentedJSON(http.StatusOK, tasks)

		}
}

func GetTaskById(tasks *[]*models.Task) gin.HandlerFunc{

	return func (c *gin.Context){
		id := c.Param("id")
		for i := range len(*tasks){
			if (*tasks)[i].ID == id {
				task := *(*tasks)[i]
				c.IndentedJSON(http.StatusOK, task)
				return
			}
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Task not found"})

	}
}

func PostTask(tasks *[]*models.Task, nextId *int) gin.HandlerFunc{
	return func (c *gin.Context){
			var task models.Task

			if err := c.BindJSON(&task); err != nil{
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
				return
			}

			task.ID = strconv.Itoa(*nextId)
			*nextId++
			*tasks = append(*tasks, &task)
			log.Println(tasks)
			c.IndentedJSON(http.StatusCreated, *((*tasks)[len(*tasks) - 1]))
		}

}


func DeleteTask(tasks *[]*models.Task) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		for i := range len(*tasks){
			if (*tasks)[i].ID == id{
				*tasks = append((*tasks)[:i], (*tasks)[i + 1:]...)
				c.IndentedJSON(http.StatusOK, gin.H{"messages" : "deleted successfully"})
				return
			}
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "no task of this id"})
	}

}

func UpdateTask(tasks *[]*models.Task) gin.HandlerFunc{
	return func (c *gin.Context){
		var task models.Task

		if err := c.BindJSON(&task); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
		} 
		
		id := c.Param("id")
		for i := range len(*tasks){
			if (*tasks)[i].ID == id{
				(*tasks)[i] = &task
				(*tasks)[i].ID = id
				c.IndentedJSON(http.StatusOK, task)
				return
			}
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "task not found"})
	}
}