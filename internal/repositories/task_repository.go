package repositories

import "github.com/Bruno07/tasks-api/internal/models"

type ITaskRepository interface {
	Save(task *models.Task) error
	Update(task *models.Task, taskId int64) error
}

type MockTaskRepository struct {
	MockSave   func(task *models.Task) error
	MockUpdate func(task *models.Task, taskId int64) error
}

func (mr *MockTaskRepository) Save(task *models.Task) error {
	return mr.MockSave(task)
}

func (mr *MockTaskRepository) Update(task *models.Task, taskId int64) error {
	return mr.MockUpdate(task, taskId)
}
