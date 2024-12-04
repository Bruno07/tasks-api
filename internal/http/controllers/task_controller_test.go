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
	c.Set("user_id", float64(1))
	c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
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
		c.Set("user_id", float64(1))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
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
		c.Set("user_id", float64(1))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
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
		c.Set("user_id", float64(1))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller := NewTaskController(*taskService)
		controller.Update(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, `{"message":"This task belongs to another user!"}`, w.Body.String())

	})

}

func TestTaskController_Find(t *testing.T) {

	gin.SetMode(gin.TestMode)

	taskRepo := &repositories.MockTaskRepository{
		MockFind: func(task *models.Task) (*models.Task, error) {

			var err error
			var result = models.Task{}

			if task.ID == 0 {
				err = errors.New("ID field is mandatory!")
			}

			var taskGroup = map[int64]models.Task{}
			taskGroup[1] = models.Task{ID: 1, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[2] = models.Task{ID: 2, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[3] = models.Task{ID: 3, Title: "Test Create Task", Description: "This is my creation test", UserID: 2}

			if task.UserID == 0 {
				if taskGroup[task.ID] != (models.Task{}) {
					result = taskGroup[task.ID]
				}

			} else {
				if taskGroup[task.ID] != (models.Task{}) && taskGroup[task.ID].UserID == task.UserID {
					result = taskGroup[task.ID]
				}
			}

			return &result, err
		},
	}

	var taskService = services.NewTaskService(taskRepo)
	var controller = NewTaskController(*taskService)

	t.Run("Must return available task to manager", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks/1", bytes.NewReader(body))
		c.Set("user_id", float64(0))
		c.AddParam("id", "1")
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.Find(c)

		var task = models.Task{}
		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, task)

	})

	t.Run("Must return empty task available to manager", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks/4", bytes.NewReader(body))
		c.Set("user_id", float64(0))
		c.AddParam("id", "4")
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.Find(c)

		var task = models.Task{}
		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, task)

	})

	t.Run("It should return an error for not finding ID", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks", bytes.NewReader(body))
		c.Set("user_id", float64(0))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.Find(c)

		var task = models.Task{}
		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Empty(t, task)

	})

	t.Run("Must return task available to technician", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks/1", bytes.NewReader(body))
		c.Set("user_id", float64(1))
		c.AddParam("id", "1")
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.Find(c)

		var task = models.Task{}
		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, task)

	})

	t.Run("Must return empty task available to technician", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks/2", bytes.NewReader(body))
		c.Set("user_id", float64(2))
		c.AddParam("id", "2")
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.Find(c)

		var task = models.Task{}
		json.Unmarshal(w.Body.Bytes(), &task)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, task)

	})

}

func TestTaskController_All(t *testing.T) {

	taskRepo := &repositories.MockTaskRepository{
		MockAll: func(task *models.Task) (*[]models.Task, error) {

			var err error
			var results = []models.Task{}

			var tasksGroup = map[int64]models.Task{}
			tasksGroup[1] = models.Task{ID: 1, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			tasksGroup[2] = models.Task{ID: 2, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			tasksGroup[3] = models.Task{ID: 3, Title: "Test Create Task", Description: "This is my creation test", UserID: 2}

			if task.UserID == 0 {
				for _, taskGroup := range tasksGroup {
					results = append(results, taskGroup)
				}
			} else {
				for _, taskGroup := range tasksGroup {
					if taskGroup.UserID == task.UserID {
						results = append(results, taskGroup)
					}
				}
			}

			return &results, err
		},
	}

	var taskService = services.NewTaskService(taskRepo)
	var controller = NewTaskController(*taskService)

	t.Run("Must return available tasks to manager", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks", bytes.NewReader(body))
		c.Set("user_id", float64(0))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.All(c)

		var tasks = []models.Task{}
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 3, len(tasks))

	})

	t.Run("Must return tasks available to technician", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks", bytes.NewReader(body))
		c.Set("user_id", float64(1))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.All(c)

		var tasks = []models.Task{}
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 2, len(tasks))

	})

	t.Run("Must return empty task available to technician", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, "/api/tasks", bytes.NewReader(body))
		c.Set("user_id", float64(3))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r

		controller.All(c)

		var tasks = []models.Task{}
		json.Unmarshal(w.Body.Bytes(), &tasks)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 0, len(tasks))

	})

}

func TestTaskService_Delete(t *testing.T) {
	taskRepo := &repositories.MockTaskRepository{
		MockDelete: func(task *models.Task) error {

			var err error

			if task.ID == 0 {
				err = errors.New("ID field is mandatory!")
			}

			var taskGroup = map[int64]models.Task{}
			taskGroup[1] = models.Task{ID: 1, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[2] = models.Task{ID: 2, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[3] = models.Task{ID: 3, Title: "Test Create Task", Description: "This is my creation test", UserID: 2}
			
			if taskGroup[task.ID] == (models.Task{}) {
				return errors.New("Task not found!")
			}

			if taskGroup[task.ID].UserID != task.UserID {
				return errors.New("This task belongs to another user!")
			}

			return err
		},
	}

	var taskService = services.NewTaskService(taskRepo)
	controller := NewTaskController(*taskService)

	t.Run("Should return no error", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodDelete, "/api/tasks/1", bytes.NewReader(body))
		c.AddParam("id", "1")
		c.Set("user_id", float64(1))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r
	
		controller.Delete(c)
	
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"message":"Task deleted successfully!"}`, w.Body.String())

	})

	t.Run("An error should be returned when an attempt is made to delete another user's task", func(t *testing.T) {

		var request = requests.TaskRequestDTO{}
		body, _ := json.Marshal(request)
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodDelete, "/api/tasks/3", bytes.NewReader(body))
		c.AddParam("id", "3")
		c.Set("user_id", float64(1))
		c.Set("permissions", []interface{}{"tasks:create", "tasks:update", "tasks:view", "tasks:delete"})
		c.Request = r
	
		controller.Delete(c)
	
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, `{"message":"This task belongs to another user!"}`, w.Body.String())

	})
}
