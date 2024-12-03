package services

import (
	"errors"
	"testing"

	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
	"github.com/stretchr/testify/assert"
)

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

	t.Run("Should return an error", func(t *testing.T) {

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
