package controllers

import (
	"log"
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(tm *data.Task_manager) gin.HandlerFunc{
		return func (c *gin.Context){
			c.IndentedJSON(http.StatusOK, tm.Tasks)

		}
}

func GetTaskById(tm *data.Task_manager) gin.HandlerFunc{

	return func (c *gin.Context){
		id := c.Param("id")
		
		task, ok := tm.Get_task(id)
		if ok{
			c.IndentedJSON(http.StatusOK, task)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Task not found"})

	}
}

func PostTask(tm *data.Task_manager) gin.HandlerFunc{
	return func (c *gin.Context){
			var task models.Task

			if err := c.BindJSON(&task); err != nil{
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
				return
			}

			task.ID = strconv.Itoa(tm.NextId)
			tm.NextId++
			tm.Tasks = append(tm.Tasks, &task)
			log.Println(tm.Tasks)
			c.IndentedJSON(http.StatusCreated, *tm.Tasks[len(tm.Tasks) - 1])
		}

}


func DeleteTask(tm *data.Task_manager) gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.Param("id")
		ok := tm.Delete_task(id)
		if ok{
			c.IndentedJSON(http.StatusOK, gin.H{"messages" : "deleted successfully"})
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "no task of this id"})
	}

}

func UpdateTask(tm *data.Task_manager) gin.HandlerFunc{
	return func (c *gin.Context){
		var task models.Task

		if err := c.BindJSON(&task); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid request"})
		} 
		
		id := c.Param("id")
		task, ok := tm.Update_task(id, task)
		if ok{
			c.IndentedJSON(http.StatusOK, task)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "task not found"})
	}
}