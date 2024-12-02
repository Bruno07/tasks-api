package controllers

import (
	"encoding/json"
	"io"
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

func (tc TaskController) Create(ctx *gin.Context) {

	if !tc.taskPolicy.Allow("CREATE", ctx) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Permission denied!",
		})

		ctx.Abort()

		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create!",
		})

		ctx.Abort()

		return
	}

	var request = requests.TaskRequest{UserID: int64(ctx.MustGet("user_id").(float64))}
	json.Unmarshal(body, &request)

	response, err := tc.taskService.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create!",
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    response,
		"message": "Successfully created!",
	})

}

func (tc TaskController) Update(ctx *gin.Context) {

	if !tc.taskPolicy.Allow("UPDATE", ctx) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Permission denied!",
		})

		ctx.Abort()

		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update!",
		})

		ctx.Abort()

		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to identify reference!",
		})

		ctx.Abort()

		return
	}

	var request = requests.TaskRequest{
		ID:     int64(taskId),
		UserID: int64(ctx.MustGet("user_id").(float64)),
	}

	json.Unmarshal(body, &request)

	response, err := tc.taskService.Update(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update!",
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    response,
		"message": "Successfully updated!",
	})

}

func (tc TaskController) Find(ctx *gin.Context) {

	if !tc.taskPolicy.Allow("VIEW", ctx) {
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

	task, err := tc.taskService.Find(requests.TaskRequest{
		ID:     int64(id),
		UserID: int64(ctx.MustGet("user_id").(float64)),
	})

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

func (tc TaskController) All(ctx *gin.Context) {

	if !tc.taskPolicy.Allow("VIEW", ctx) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Permission denied!",
		})

		ctx.Abort()

		return
	}

	task, err := tc.taskService.All(requests.TaskRequest{
		UserID: int64(ctx.MustGet("user_id").(float64)),
	})

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
