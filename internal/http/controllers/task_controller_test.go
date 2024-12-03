package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTaskController_Create(t *testing.T) {

	taskRepo := &repositories.MockTaskRepository{
		MockSave: func(task *models.Task) (err error) {
			err = task.Validate()
			return err
			
		},
	}

	var taskService = services.NewTaskService(taskRepo)

	var request = requests.TaskRequestDTO{
		Title:       "Test create task",
		Description: "This is my create test",
		User:        requests.UserRequestDTO{ID: 1},
	}

	body, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	r, _ := http.NewRequest(http.MethodPost, "/api/tasks", bytes.NewReader(body))
	c.Request = r

	controller := NewTaskController(*taskService)
	controller.Create(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{"message":"Task registered successfully!"}`, w.Body.String())
	
}
