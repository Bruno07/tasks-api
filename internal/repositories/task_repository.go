package repositories

import (
	"errors"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type ITaskRepository interface {
	Save(request requests.TaskRequest) (models.Task, error)
	Update(request requests.TaskRequest) (models.Task, error)
	Find(request requests.TaskRequest) (models.Task, int64, error)
	All(request requests.TaskRequest) ([]models.Task, error)
	Delete(taskId int64) bool
}

type TaskRepository struct{}

type TaskMockRepository struct{}

// Persists task in database
func (tr TaskRepository) Save(request requests.TaskRequest) (models.Task, error) {

	task := models.Task{
		Summary: request.Summary,
		UserID:  request.UserID,
	}

	err := db.GetInstance().Create(&task).Error

	return task, err

}

// Update task in database
func (tr TaskRepository) Update(request requests.TaskRequest) (models.Task, error) {

	task := models.Task{}
	query := db.GetInstance().Model(&task)

	if request.ID != 0 {
		query.Where("id = ?", request.ID)
	}

	if request.UserID != 0 {
		query.Where("user_id = ?", request.UserID)
	}

	result := query.Updates(map[string]interface{}{
		"summary": request.Summary,
	})

	return task, result.Error

}

// Find user through filter
func (tr TaskRepository) Find(request requests.TaskRequest) (models.Task, int64, error) {

	var task models.Task

	query := db.GetInstance().Model(&task)

	if request.ID != 0 {
		query.Where("id = ?", request.ID)
	}

	if request.UserID != 0 {
		query.Where("user_id = ?", request.UserID)
	}

	result := query.Find(&task)

	return task, result.RowsAffected, result.Error

}

// List all tasks through filter
func (tr *TaskRepository) All(request requests.TaskRequest) ([]models.Task, error) {

	var task = models.Task{}
	var tasks []models.Task

	query := db.GetInstance().Model(&task)

	if request.ID != 0 {
		query.Where("id = ?", request.ID)
	}

	if request.UserID != 0 {
		query.Where("user_id = ?", request.UserID)
	}

	err := query.Find(&tasks).Error

	return tasks, err

}

func (tr TaskRepository) Delete(taskId int64) bool {

	var isDeleted bool = false

	result := db.GetInstance().Delete(&models.Task{}, taskId)

	if result.RowsAffected > 0 {
		isDeleted = true
	}

	return isDeleted

}

func (tmr *TaskMockRepository) Save(request requests.TaskRequest) (models.Task, error) {
	var err error
	var task = models.Task{
		Summary: request.Summary,
		PerformedDate: &request.PerformedDate,
	}

	if request.Summary == "" {
		err = errors.New("Failed to create task!")
	}

	return task, err
}

func (tmr *TaskMockRepository) Update(request requests.TaskRequest) (models.Task, error) {

	var err error
	var task = models.Task{}

	if request.ID == 0 {
		err = errors.New("Failed to update task!")
	}

	if request.ID > 0 {
		task.Summary = request.Summary
		task.PerformedDate = &request.PerformedDate
	}

	return task, err
}

func (tmr *TaskMockRepository) Find(request requests.TaskRequest) (models.Task, int64, error) {

	var err error
	var rowsAffected int64 = 0
	task := models.Task{}

	if request.Summary == "Lorem" {
		err = errors.New("Invalid summary")
	}

	if request.ID <= 3 {
		task = models.Task{
			ID:            request.ID,
			Summary:       request.Summary,
			PerformedDate: &request.PerformedDate,
		}

		rowsAffected = 1
	}

	return task, rowsAffected, err

}

func (tmr *TaskMockRepository) All(request requests.TaskRequest) ([]models.Task, error) {

	var err error
	var tasks = make([]models.Task, 0)

	if request.ID <= 3 {
		task := models.Task{
			ID:            request.ID,
			Summary:       request.Summary,
			PerformedDate: &request.PerformedDate,
		}

		tasks = append(tasks, task)

	} else {
		if request.Summary == "Lorem" {
			err = errors.New("Invalid summary")
		}
	}

	return tasks, err

}

func (tmr *TaskMockRepository) Delete(taskId int64) bool {

	var isDeleted bool = false

	if taskId <= 3 {
		isDeleted = true
	}

	return isDeleted

}
