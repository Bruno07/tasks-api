package services

import (
	"errors"
	"testing"

	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
	"github.com/stretchr/testify/assert"
)

// Task creation test
func TestTaskService_Create(t *testing.T) {

	taskRepo := &repositories.MockTaskRepository{
		MockSave: func(task *models.Task) (err error) {
			if task.Title == "" {
				err = errors.New("Title field is mandatory!")
			}

			if task.Description == "" {
				err = errors.New("Description field is mandatory!")
			}

			return err
		},
	}

	t.Run("Must create a task", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		err := service.Create(&requests.TaskRequestDTO{
			Title:       "Test Create Task",
			Description: "This is my creation test",
			User: requests.UserRequestDTO{
				ID: 1,
			},
		})

		assert.NoError(t, err)

	})

	t.Run("It should return a validation error", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		err := service.Create(&requests.TaskRequestDTO{
			Title:       "",
			Description: "This is my creation test",
			User: requests.UserRequestDTO{
				ID: 1,
			},
		})

		assert.Error(t, err)

	})

}

// Task Update Test
func TestTaskService_Update(t *testing.T) {

	taskRepo := &repositories.MockTaskRepository{
		MockUpdate: func(task *models.Task, taskId int64) (err error) {
			if task.Title == "" {
				err = errors.New("Title field is mandatory!")
			}

			if task.Description == "" {
				err = errors.New("Description field is mandatory!")
			}

			var taskGroup = map[int64]models.Task{}
			taskGroup[1] = models.Task{ID: 1, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[2] = models.Task{ID: 2, Title: "Test Create Task", Description: "This is my creation test", UserID: 1}
			taskGroup[3] = models.Task{ID: 3, Title: "Test Create Task", Description: "This is my creation test", UserID: 2}

			if taskGroup[taskId].UserID != task.UserID {
				err = errors.New("This task belongs to another user!")
			}

			return err
		},
	}

	t.Run("Must update a task", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		err := service.Update(&requests.TaskRequestDTO{
			Title:       "Test Update Task",
			Description: "This is my update test",
			User: requests.UserRequestDTO{
				ID: 1,
			},
		}, 2)

		assert.NoError(t, err)

	})

	t.Run("It should return a validation error", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		err := service.Update(&requests.TaskRequestDTO{
			Title:       "Test Update Task",
			Description: "",
			User: requests.UserRequestDTO{
				ID: 1,
			},
		}, 1)

		assert.Error(t, err)

	})

	t.Run("Should return an error when trying to change a task for another user", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		err := service.Update(&requests.TaskRequestDTO{
			Title:       "Test Update Task",
			Description: "This is my creation test",
			User: requests.UserRequestDTO{
				ID: 2,
			},
		}, 1)

		assert.Error(t, err)

	})
}
