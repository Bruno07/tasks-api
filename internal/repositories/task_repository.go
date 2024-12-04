package repositories

import (
	"github.com/Bruno07/tasks-api/internal/models"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	Save(task *models.Task) error
	Update(task *models.Task, taskId int64) error
	Find(task *models.Task) (*models.Task, error)
	All(task *models.Task) (*[]models.Task, error)
	Delete(task *models.Task) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

type MockTaskRepository struct {
	MockSave   func(task *models.Task) error
	MockUpdate func(task *models.Task, taskId int64) error
	MockFind   func(task *models.Task) (*models.Task, error)
	MockAll    func(task *models.Task) (*[]models.Task, error)
	MockDelete func(task *models.Task) error
}

// Save a task
func (mr *TaskRepository) Save(task *models.Task) error {
	return mr.db.Create(&task).Error
}

// Update a task
func (mr *TaskRepository) Update(task *models.Task, taskId int64) error {
	query := mr.db.Model(&task).Where("id = ?", taskId)

	if task.UserID != 0 {
		query.Where("user_id = ?", task.UserID)
	}

	result := query.Updates(map[string]interface{}{
		"title": task.Title,
		"description": task.Description,
	})

	return result.Error
}

// Find a task
func (mr *TaskRepository) Find(task *models.Task) (*models.Task, error) {

	query := mr.db.Model(&task) 
	
	if task.UserID != 0 {
		query.Where("user_id = ?", task.UserID)
	}
	
	result := query.Find(&task)

	return task, result.Error

}

func (mr *TaskRepository) All(task *models.Task) (*[]models.Task, error) {

	var tasks = []models.Task{}

	query := mr.db.Model(&task)

	if task.UserID != 0 {
		query.Where("user_id = ?", task.UserID)
	}

	result := query.Find(&tasks)

	return &tasks, result.Error
}

// Delete a task
func (mr *TaskRepository) Delete(task *models.Task) error {
	return mr.db.Delete(&task, task.ID).Error
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
func (mr *MockTaskRepository) Delete(task *models.Task) error {
	return mr.MockDelete(task)
}
