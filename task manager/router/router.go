package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

func Run(){
	router := gin.Default()

	router.GET("api/tasks", controllers.GetAllTasks)
	router.GET("api/tasks/:id")
	router.POST("api/tasks")
	router.PUT("api/tasks/:id")
	router.DELETE("api/tasks/:id")

}