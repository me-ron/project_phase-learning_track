package routes

import (
	"task_manager/database"
	"task_manager/delivery/controllers"
	"task_manager/infrastructure"
	"task_manager/repository"
	"task_manager/useCase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartTaskRoutes(db *mongo.Database, router *gin.Engine){
	collection := &database.MongoCollection{Collection: db.Collection("tasks")}
	task_repo := repository.NewTaskRepo(collection)
	task_usecase := useCase.NewTaskUC(task_repo)
	loggedIn := router.Group("")
	loggedIn.Use(infrastructure.AuthMiddleware)

	loggedIn.GET("api/tasks", infrastructure.RoleBasedAuth(false), controllers.GetAllTasks(task_usecase))
	loggedIn.GET("api/tasks/:id", infrastructure.RoleBasedAuth(false), controllers.GetTaskById(task_usecase))
	loggedIn.POST("api/tasks", infrastructure.RoleBasedAuth(false), controllers.PostTask(task_usecase))
	loggedIn.PUT("api/tasks/:id", infrastructure.RoleBasedAuth(false), controllers.UpdateTask(task_usecase))
	loggedIn.DELETE("api/tasks/:id", infrastructure.RoleBasedAuth(false), controllers.DeleteTask(task_usecase))
}