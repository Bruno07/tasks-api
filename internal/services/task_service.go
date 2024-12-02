package services

import (
	"encoding/json"
	"errors"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/http/responses"
	"github.com/Bruno07/tasks-api/internal/infra/queue"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
)

type TaskService struct {
	taskRepo repositories.ITaskRepository
	userRepo repositories.IUserRepository
}

func NewTaskService(
	taskRepo repositories.ITaskRepository,
	userRepo repositories.IUserRepository,
) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

// Create tasks and notify manager
func (ts TaskService) Create(request requests.TaskRequest) (responses.TaskResponse, error) {

	taskCreated, err := ts.taskRepo.Save(request)
	if err != nil {
		return responses.TaskResponse{}, err
	}

	notification := models.Notification{Payload: "User created a new task", UserID: request.UserID}
	body, _ := json.Marshal(notification)
	ch := queue.GetInstanceQueue()
	queue.Notify(body, "notify_ex", "", ch)

	user, _ := ts.userRepo.Find(taskCreated.UserID)

	response := responses.TaskResponse{
		ID:            taskCreated.ID,
		Summary:       taskCreated.Summary,
		PerformedDate: taskCreated.PerformedDate,
		User: responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil
}

// List a task created by the user
func (ts TaskService) Find(request requests.TaskRequest) (responses.TaskResponse, error) {

	task, rowsAffected, err := ts.taskRepo.Find(request)

	if err != nil {
		return responses.TaskResponse{}, err
	}

	if rowsAffected == 0 {
		return responses.TaskResponse{}, nil
	}

	user, _ := ts.userRepo.Find(task.UserID)

	response := responses.TaskResponse{
		ID:            task.ID,
		Summary:       task.Summary,
		PerformedDate: task.PerformedDate,
		User: responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil

}

// List all tasks created by the user
func (ts *TaskService) All(request requests.TaskRequest) ([]responses.TaskResponse, error) {

	tasks, err := ts.taskRepo.All(request)

	if err != nil {
		return nil, err
	}

	var tasksReponse = []responses.TaskResponse{}
	for _, task := range *tasks {

		user, _ := ts.userRepo.Find(task.UserID)

		var response = responses.TaskResponse{
			ID:            task.ID,
			Summary:       task.Summary,
			PerformedDate: task.PerformedDate,
			User: responses.UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		}

		tasksReponse = append(tasksReponse, response)
	}

	return tasksReponse, nil
}

// Update user created task
func (ts TaskService) Update(request requests.TaskRequest) (responses.TaskResponse, error) {

	taskUpdated, err := ts.taskRepo.Update(request)

	if err != nil {
		return responses.TaskResponse{}, err
	}

	notification := models.Notification{Payload: "User updated a new task", UserID: request.UserID}
	body, _ := json.Marshal(notification)
	ch := queue.GetInstanceQueue()
	queue.Notify(body, "notify_ex", "", ch)

	user, _ := ts.userRepo.Find(taskUpdated.UserID)

	response := responses.TaskResponse{
		ID:            taskUpdated.ID,
		Summary:       taskUpdated.Summary,
		PerformedDate: taskUpdated.PerformedDate,
		User: responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil

}

func (ts TaskService) Delete(request requests.TaskRequest) (err error) {

	if ts.taskRepo.Delete(request.ID) == false {
		err = errors.New("Failed to delete task!")
	}

	return err
}
