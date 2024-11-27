package routes

import (
	"github.com/Bruno07/tasks-api/internal/http/controllers"
	"github.com/Bruno07/tasks-api/internal/http/middlewares"
	"github.com/gin-gonic/gin"
)


func InitRoutes() *gin.Engine {

	router := gin.New()

	router.SetTrustedProxies([]string{"0.0.0.0"})
	
	router.POST("/login", controllers.NewAuthController().Login)

	auth := router.Group("auth")
	auth.Use(middlewares.AuthMiddleware())
	auth.POST("/", controllers.NewUserController().Create)

	return router

}

