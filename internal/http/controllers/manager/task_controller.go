package manager

import (
	"net/http"
	"strconv"

	"github.com/Bruno07/tasks-api/internal/http/policies"
	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService services.TaskService
	TaskPolicy  policies.TaskPolicy
}

func NewTaskController() *TaskController {
	return &TaskController{
		TaskService: services.TaskService{},
		TaskPolicy:  policies.TaskPolicy{},
	}
}

func (tc TaskController) All(ctx *gin.Context) {

	if !tc.TaskPolicy.Allow("VIEW", ctx) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Permission denied!",
		})

		ctx.Abort()

		return
	}

	task, err := tc.TaskService.All(requests.TaskRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get task listing!",
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}

func (tc TaskController) Delete(ctx *gin.Context) {

	if !tc.TaskPolicy.Allow("DELETE", ctx) {
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

	err = tc.TaskService.Delete(requests.TaskRequest{ID: int64(id)})

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
