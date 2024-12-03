package repositories

import "github.com/Bruno07/tasks-api/internal/models"

type ITaskRepository interface {
	Save(task *models.Task) error
}

type MockTaskRepository struct {
	MockSave func(task *models.Task) error
}

func (mr *MockTaskRepository) Save(task *models.Task) error {
	return mr.MockSave(task)
}
