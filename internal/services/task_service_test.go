package services

import (
	"testing"
	"time"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestTaskService_Create(t *testing.T) {

	t.Run("Should return a task", func(t *testing.T) {

		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            1,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		response, err := taskService.Create(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

	})

	t.Run("Should return an error", func(t *testing.T) {

		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            0,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		_, err := taskService.Update(request)

		assert.Error(t, err)

	})

}

func TestTaskService_Update(t *testing.T) {

	t.Run("Should return a task", func(t *testing.T) {

		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            1,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		response, err := taskService.Update(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

	})

	t.Run("Should return an error", func(t *testing.T) {

		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            0,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		_, err := taskService.Update(request)

		assert.Error(t, err)

	})

}

func TestTaskService_Find(t *testing.T) {
	t.Run("Should return a task", func(t *testing.T) {

		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            1,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		response, err := taskService.Find(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

	})

	t.Run("Must return an empty task", func(t *testing.T) {

		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            4,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		response, err := taskService.Find(request)

		assert.NoError(t, err)
		assert.Empty(t, response)

	})

	t.Run("Should return an error", func(t *testing.T) {

		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            4,
			Summary:       "Lorem",
			PerformedDate: time.Now(),
		}

		_, err := taskService.Find(request)

		assert.Error(t, err)

	})
}

func TestTaskService_All(t *testing.T) {

	taskService := NewTaskService(
		&repositories.TaskMockRepository{},
		&repositories.UserMockRepository{},
	)

	t.Run("Should return a list of tasks without errors", func(t *testing.T) {

		var request = requests.TaskRequest{
			ID:            1,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		response, err := taskService.All(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

	})

	t.Run("Must return an empty list", func(t *testing.T) {

		var request = requests.TaskRequest{ID: 4}

		response, err := taskService.All(request)

		assert.NoError(t, err)
		assert.Empty(t, response)

	})

}

func TestTaskService_Delete(t *testing.T) {

	t.Run("It should return without an error", func(t *testing.T) {
		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		err := taskService.Delete(requests.TaskRequest{
			ID: 1,
		})

		assert.NoError(t, err)
	})

	t.Run("It should return with an error", func(t *testing.T) {
		taskService := NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		err := taskService.Delete(requests.TaskRequest{
			ID: 4,
		})

		assert.Error(t, err)
	})
}
