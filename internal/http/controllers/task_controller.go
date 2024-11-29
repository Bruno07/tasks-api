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
	TaskService services.TaskService
	TaskPolicy  policies.TaskPolicy
}

func NewTaskController() *TaskController {
	return &TaskController{
		TaskService: services.TaskService{},
		TaskPolicy:  policies.TaskPolicy{},
	}
}

func (tc TaskController) Create(ctx *gin.Context) {

	if !tc.TaskPolicy.Allow("CREATE", ctx) {
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

	response, err := tc.TaskService.Create(request)
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

	if !tc.TaskPolicy.Allow("UPDATE", ctx) {
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

	response, err := tc.TaskService.Update(request)
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

	if !tc.TaskPolicy.Allow("VIEW", ctx) {
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

	task, err := tc.TaskService.Find(requests.TaskRequest{
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

	if !tc.TaskPolicy.Allow("VIEW", ctx) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Permission denied!",
		})

		ctx.Abort()

		return
	}

	task, err := tc.TaskService.All(requests.TaskRequest{
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
