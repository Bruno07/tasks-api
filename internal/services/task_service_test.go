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

func TestTaskService_Find(t *testing.T) {
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

	t.Run("Must return available task to manager", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		result, err := service.Find(&requests.TaskRequestDTO{
			ID: 2,
			User: requests.UserRequestDTO{ID: 0},
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, result)

	})

	t.Run("Must return empty task available to manager", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		result, err := service.Find(&requests.TaskRequestDTO{
			ID: 4,
			User: requests.UserRequestDTO{ID: 0},
		})

		assert.NoError(t, err)
		assert.Empty(t, result)

	})

	t.Run("It should return an error for not finding ID", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		result, err := service.Find(&requests.TaskRequestDTO{
			ID: 0,
			User: requests.UserRequestDTO{ID: 0},
		})

		assert.Error(t, err)
		assert.Empty(t, result)

	})

	t.Run("Must return task available to technician", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		result, err := service.Find(&requests.TaskRequestDTO{
			ID: 2,
			User: requests.UserRequestDTO{ID: 1},
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, result)

	})

	t.Run("Must return empty task available to technician", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		result, err := service.Find(&requests.TaskRequestDTO{
			ID: 2,
			User: requests.UserRequestDTO{ID: 2},
		})

		assert.NoError(t, err)
		assert.Empty(t, result)

	})
}

func TestTaskService_All(t *testing.T) {
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

	t.Run("Must return available tasks to manager", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		results, err := service.GetAll(&requests.TaskRequestDTO{})

		assert.NoError(t, err)
		assert.Equal(t, 3, len(*results))

	})

	t.Run("Must return tasks available to technician", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		results, err := service.GetAll(&requests.TaskRequestDTO{
			User: requests.UserRequestDTO{ID: 1},
		})

		assert.NoError(t, err)
		assert.Equal(t, 2, len(*results))

	})

	t.Run("Must return empty task available to technician", func(t *testing.T) {

		service := NewTaskService(taskRepo)
		results, err := service.GetAll(&requests.TaskRequestDTO{
			User: requests.UserRequestDTO{ID: 3},
		})

		assert.NoError(t, err)
		assert.Equal(t, 0, len(*results))

	})
}