package router

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func Run(task_manager *data.Taskmanager) {
	router := gin.New()
	loggedIn := router.Group("")
	loggedIn.Use(middleware.AuthMiddleware)

	loggedIn.GET("api/tasks", middleware.RoleBasedAuth(false), controllers.GetAllTasks(task_manager))
	loggedIn.GET("api/tasks/:id", middleware.RoleBasedAuth(false), controllers.GetTaskById(task_manager))
	loggedIn.POST("api/tasks", middleware.RoleBasedAuth(false), controllers.PostTask(task_manager))
	loggedIn.PUT("api/tasks/:id", middleware.RoleBasedAuth(false), controllers.UpdateTask(task_manager))
	loggedIn.DELETE("api/tasks/:id", middleware.RoleBasedAuth(false), controllers.DeleteTask(task_manager))
	router.POST("api/register", controllers.Register(task_manager))
	router.POST("api/login", controllers.Login(task_manager))

	router.Run("localhost:8080")

}
