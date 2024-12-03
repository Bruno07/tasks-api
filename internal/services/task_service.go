package services

import (
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
)

type TaskService struct {
	taskRepo repositories.ITaskRepository
}

func NewTaskService(taskRepo repositories.ITaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

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
