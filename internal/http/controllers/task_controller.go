package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Bruno07/tasks-api/internal/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (tc *TaskController) Create(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)

	var request = &requests.TaskRequestDTO{}
	if err := json.Unmarshal(body, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create task!",
		})

		c.Abort()
		
		return
	}

	if err := tc.taskService.Create(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task registered successfully!",
	})

}