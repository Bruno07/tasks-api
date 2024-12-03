package services

import (
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
)

type TaskService struct {
	taskRepo repositories.ITaskRepository
}

// Create an instance of the task service
func NewTaskService(taskRepo repositories.ITaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

// Create a task
func (ts *TaskService) Create(request *requests.TaskRequestDTO) (err error) {

	var task = models.Task{
		Title:       request.Title,
		Description: request.Description,
		UserID:      request.User.ID,
	}

	err = ts.taskRepo.Save(&task)
	if err != nil {
		return err
	}

	return err

}

// Update a task
func (ts *TaskService) Update(request *requests.TaskRequestDTO, taskId int64) (err error) {

	var task = models.Task{
		Title:       request.Title,
		Description: request.Description,
		UserID:      request.User.ID,
	}

	err = ts.taskRepo.Update(&task, taskId)
	if err != nil {
		return err
	}

	return err

}

// Find a task
func (ts *TaskService) Find(request *requests.TaskRequestDTO) (*models.Task, error) {

	var task = models.Task{
		ID:     request.ID,
		UserID: request.User.ID,
	}

	result, err := ts.taskRepo.Find(&task)
	if err != nil {
		return nil, err
	}

	return result, err

}

// Get all tasks
func (ts *TaskService) GetAll(request *requests.TaskRequestDTO) (*[]models.Task, error) {

	var task = models.Task{
		UserID: request.User.ID,
	}

	results, err := ts.taskRepo.All(&task)
	if err != nil {
		return nil, err
	}

	return results, err

}

func (ts *TaskService) Delete(taskId int64) (err error) {

	err = ts.taskRepo.Delete(taskId)
	if err != nil {
		return err
	}

	return err

}
