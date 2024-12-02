package routes

import (
	"github.com/Bruno07/tasks-api/internal/http/controllers"
	"github.com/Bruno07/tasks-api/internal/http/controllers/manager"
	"github.com/Bruno07/tasks-api/internal/http/middlewares"
	"github.com/Bruno07/tasks-api/internal/http/policies"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)

// Start route configuration
func InitRoutes() *gin.Engine {

	router := gin.New()

	router.SetTrustedProxies([]string{"0.0.0.0"})

	router.POST("/login", controllers.NewAuthController().Login)

	// Protected route group
	auth := router.Group("api")
	auth.Use(middlewares.AuthMiddleware())

	auth.POST("/register", manager.NewUserController().Create)

	auth.POST("/tasks", controllers.NewTaskController().Create)
	auth.PUT("/tasks/:id", controllers.NewTaskController().Update)
	auth.GET("/tasks", controllers.NewTaskController().All)
	auth.GET("/tasks/:id", controllers.NewTaskController().Find)

	taskService := services.NewTaskService(
		&repositories.TaskRepository{},
		&repositories.UserRepository{},
	)
	
	taskController := manager.NewTaskController(
		*taskService,
		policies.TaskPolicy{},
	)

	auth.GET("/manager/tasks", taskController.All)
	auth.DELETE("/manager/tasks/:id", taskController.Delete)

	return router

}
