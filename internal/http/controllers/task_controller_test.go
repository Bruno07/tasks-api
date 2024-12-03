package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
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

	gin.SetMode(gin.TestMode)

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

func TestTaskController_Update(t *testing.T) {

	gin.SetMode(gin.TestMode)

	taskRepo := &repositories.MockTaskRepository{
		MockUpdate: func(task *models.Task, taskId int64) (err error) {
			err = task.Validate()
			if err != nil {
				return err
			}

			var taskGroup = map[int64]models.Task{}
			taskGroup[1] = models.Task{ID: 1, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[2] = models.Task{ID: 2, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[3] = models.Task{ID: 3, Title: "Test Create Task", Description: "This is my creation test", UserID: 2}

			if taskGroup[taskId] == (models.Task{}) {
				return errors.New("Task not found!")
			}

			if taskGroup[taskId].UserID != task.UserID {
				return errors.New("This task belongs to another user!")
			}

			return err

		},
	}

	var taskService = services.NewTaskService(taskRepo)

	t.Run("Should return status 201 and success message", func(t *testing.T) {

		var request = requests.TaskRequestDTO{
			Title:       "Test update task",
			Description: "This is my update test",
			User:        requests.UserRequestDTO{ID: 1},
		}
	
		body, _ := json.Marshal(request)
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodPut, "/api/tasks/1", bytes.NewReader(body))
		c.AddParam("id", "1")
		c.Request = r
	
		controller := NewTaskController(*taskService)
		controller.Update(c)
	
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, `{"message":"Task updated successfully!"}`, w.Body.String())

	})

	t.Run("Should return an error when not finding the task", func(t *testing.T) {

		var request = requests.TaskRequestDTO{
			Title:       "Test update task",
			Description: "This is my update test",
			User:        requests.UserRequestDTO{ID: 1},
		}
	
		body, _ := json.Marshal(request)
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodPut, "/api/tasks/4", bytes.NewReader(body))
		c.AddParam("id", "4")
		c.Request = r
	
		controller := NewTaskController(*taskService)
		controller.Update(c)
	
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, `{"message":"Task not found!"}`, w.Body.String())

	})

	t.Run("Should return an error when trying to update a task from another user", func(t *testing.T) {

		var request = requests.TaskRequestDTO{
			Title:       "Test update task",
			Description: "This is my update test",
			User:        requests.UserRequestDTO{ID: 2},
		}
	
		body, _ := json.Marshal(request)
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodPut, "/api/tasks/1", bytes.NewReader(body))
		c.AddParam("id", "1")
		c.Request = r
	
		controller := NewTaskController(*taskService)
		controller.Update(c)
	
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, `{"message":"This task belongs to another user!"}`, w.Body.String())

	})


}
