package routes

import (
	"github.com/Bruno07/tasks-api/internal/http/controllers"
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)

var (
	taskController *controllers.TaskController
)

func init() {

	var taskRepo = repositories.NewTaskRepository(db.GetInstance())
	var taskService = services.NewTaskService(taskRepo)
	taskController = controllers.NewTaskController(*taskService)

}

func LoadRoutes() *gin.Engine {

	router := gin.Default()

	router.POST("/tasks", taskController.Create)

	return router
}
