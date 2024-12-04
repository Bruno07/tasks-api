package routes

import (
	"github.com/Bruno07/tasks-api/internal/http/controllers"
	"github.com/Bruno07/tasks-api/internal/http/middleware"
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)


func LoadRoutes() *gin.Engine {

	router := gin.Default()

	var userRepo = repositories.NewUserRepository(db.GetInstance())
	var authService = services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(*authService)

	router.POST("/login", authController.Login)

	var taskRepo = repositories.NewTaskRepository(db.GetInstance())
	var taskService = services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(*taskService)

	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	auth.POST("/tasks", taskController.Create)
	auth.GET("/tasks", taskController.All)
	auth.GET("/tasks/:id", taskController.Find)
	auth.PUT("/tasks/:id", taskController.Update)
	auth.DELETE("/tasks/:id", taskController.Delete)

	return router
}
