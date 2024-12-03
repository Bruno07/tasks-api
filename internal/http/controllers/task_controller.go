package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

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

func (tc *TaskController) Update(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)

	var request = &requests.TaskRequestDTO{}
	if err := json.Unmarshal(body, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update task!",
		})

		c.Abort()

		return
	}

	taskId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := tc.taskService.Update(request, taskId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task updated successfully!",
	})

}

func (tc *TaskController) Find(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)

	var request = requests.TaskRequestDTO{}
	if err := json.Unmarshal(body, &request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to find task!",
		})

		c.Abort()

		return
	}

	userId := c.MustGet("user_id").(int64)
	taskId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	request.ID = taskId
	request.User.ID = userId

	result, err := tc.taskService.Find(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result)

}

func (tc *TaskController) All(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)

	var request = requests.TaskRequestDTO{}
	if err := json.Unmarshal(body, &request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to find task!",
		})

		c.Abort()

		return
	}

	request.User.ID = c.MustGet("user_id").(int64)
	
	results, err := tc.taskService.GetAll(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		c.Abort()

		return
	}

	c.JSON(http.StatusOK, results)

}
