package manager

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Bruno07/tasks-api/internal/http/policies"
	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService services.TaskService
	taskPolicy  policies.TaskPolicy
}

func NewTaskController(
	taskService services.TaskService,
	taskPolicy policies.TaskPolicy,
) *TaskController {
	return &TaskController{
		taskService: taskService,
		taskPolicy:  taskPolicy,
	}
}

func (tc TaskController) All(ctx *gin.Context) {

	if !tc.taskPolicy.Allow("VIEW", ctx) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Permission denied!",
		})

		ctx.Abort()

		return
	}

	task, err := tc.taskService.All(requests.TaskRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get task listing!",
		})

		ctx.Abort()

		return
	}

	fmt.Println(task)

	ctx.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}

func (tc TaskController) Delete(ctx *gin.Context) {

	if !tc.taskPolicy.Allow("DELETE", ctx) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Permission denied!",
		})

		ctx.Abort()

		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to identify reference!",
		})

		ctx.Abort()

		return
	}

	err = tc.taskService.Delete(requests.TaskRequest{ID: int64(id)})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully!",
	})

}
