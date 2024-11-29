package manager

import (
	"testing"
	"time"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestTaskController_All(t *testing.T) {

	t.Run("Should return a list of tasks without errors", func(t *testing.T) {

		taskService := services.NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            1,
			Summary:       "Lorem Ipsum",
			PerformedDate: time.Now(),
		}

		response, err := taskService.All(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

	})

	t.Run("Must return an empty list without errors", func(t *testing.T) {

		taskService := services.NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{ID: 4}

		response, err := taskService.All(request)

		assert.NoError(t, err)
		assert.Empty(t, response)

	})
	t.Run("Must return an empty list with errors", func(t *testing.T) {

		taskService := services.NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		var request = requests.TaskRequest{
			ID:            4,
			Summary:       "Lorem",
			PerformedDate: time.Now(),
		}

		response, err := taskService.All(request)

		assert.Error(t, err)
		assert.Empty(t, response)

	})

}

func TestTaskController_Delete(t *testing.T) {

	t.Run("It should return without an error", func(t *testing.T) {
		taskService := services.NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		err := taskService.Delete(requests.TaskRequest{
			ID: 1,
		})
	
		assert.NoError(t, err)
	})

	t.Run("It should return with an error", func(t *testing.T) {
		taskService := services.NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		err := taskService.Delete(requests.TaskRequest{
			ID: 4,
		})
	
		assert.Error(t, err)
	})
}
