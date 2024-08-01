package router

import (
	"task_manager/controllers"
	"task_manager/data"
	"github.com/gin-gonic/gin"
)

func Run(task_manager data.Task_manager){
	router := gin.Default()

	router.GET("api/tasks", controllers.GetAllTasks(&task_manager))
	router.GET("api/tasks/:id", controllers.GetTaskById(&task_manager))
	router.POST("api/tasks", controllers.PostTask(&task_manager))
	router.PUT("api/tasks/:id", controllers.UpdateTask(&task_manager))
	router.DELETE("api/tasks/:id", controllers.DeleteTask(&task_manager))

	router.Run("localhost:8080")

}