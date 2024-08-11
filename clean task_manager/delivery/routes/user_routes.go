package routes

import (
	"log"
	"task_manager/delivery/controllers"
	"task_manager/infrastructure"
	"task_manager/repository"
	"task_manager/useCase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartUserRoutes(db *mongo.Database, router *gin.Engine){
	user_repo, err := repository.NewUserRepo(db, "users")
	if err != nil{
		log.Panic(err.Error())
	}

	pass_s := infrastructure.PasswordS{}
	token_s := infrastructure.JwtService{}
	user_usecase := useCase.NewUserUC(user_repo, &pass_s, token_s)

	router.POST("api/register", controllers.Register(user_usecase))
	router.POST("api/login", controllers.Login(user_usecase))

	loggedIn := router.Group("")
	loggedIn.Use(infrastructure.AuthMiddleware)

	loggedIn.GET("api/users", infrastructure.RoleBasedAuth(true), controllers.GetAllUsers(user_usecase))
	loggedIn.GET("api/users/:id", infrastructure.RoleBasedAuth(false), controllers.GetUserById(user_usecase))
	loggedIn.PUT("api/users/:id", infrastructure.RoleBasedAuth(false), controllers.UpdateUser(user_usecase))
	loggedIn.DELETE("api/users/:id", infrastructure.RoleBasedAuth(false), controllers.DeleteUser(user_usecase))
	loggedIn.PUT("api/users/:id/make-admin", infrastructure.RoleBasedAuth(true), controllers.MakeAdmin(user_usecase))
}