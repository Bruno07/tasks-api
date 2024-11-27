package routes

import (
	"github.com/Bruno07/tasks-api/internal/http/controllers"
	"github.com/gin-gonic/gin"
)


func InitRoutes() *gin.Engine {

	router := gin.New()

	router.SetTrustedProxies([]string{"0.0.0.0"})

	router.POST("/", controllers.NewUserController().Create)

	return router

}

