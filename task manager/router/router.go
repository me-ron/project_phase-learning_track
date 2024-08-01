package router

import (
	"task_manager/controllers"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

var tasks []*models.Task

func Run(){
	router := gin.Default()

	router.GET("api/tasks", controllers.GetAllTasks(tasks))
	router.GET("api/tasks/:id", controllers.GetTaskById(tasks))
	router.POST("api/tasks", controllers.PostTask(tasks))
	router.PUT("api/tasks/:id", controllers.UpdateTask(tasks))
	router.DELETE("api/tasks/:id", controllers.DeleteTask(tasks))

	router.Run("localhost:8080/")

}