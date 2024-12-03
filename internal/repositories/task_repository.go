package repositories

import "github.com/Bruno07/tasks-api/internal/models"

type ITaskRepository interface {
	Save(task *models.Task) error
	Update(task *models.Task, taskId int64) error
	Find(task *models.Task) (*models.Task, error)
	All(task *models.Task) (*[]models.Task, error)
	Delete(taskId int64) error
}

type MockTaskRepository struct {
	MockSave   func(task *models.Task) error
	MockUpdate func(task *models.Task, taskId int64) error
	MockFind   func(task *models.Task) (*models.Task, error)
	MockAll    func(task *models.Task) (*[]models.Task, error)
	MockDelete func(taskId int64) error
}

// Save a task
func (mr *MockTaskRepository) Save(task *models.Task) error {
	return mr.MockSave(task)
}

// Update a task
func (mr *MockTaskRepository) Update(task *models.Task, taskId int64) error {
	return mr.MockUpdate(task, taskId)
}

// Find a task
func (mr *MockTaskRepository) Find(task *models.Task) (*models.Task, error) {
	return mr.MockFind(task)
}

func (mr *MockTaskRepository) All(task *models.Task) (*[]models.Task, error) {
	return mr.MockAll(task)
}

// Delete a task
func (mr *MockTaskRepository) Delete(taskId int64) error {
	return mr.MockDelete(taskId)
}
