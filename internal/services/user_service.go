package services

import (
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) *UserService{
	return &UserService{
		userRepo: userRepo,
	}
}

// Create a user
func (s *UserService) Create(request *requests.UserRequestDTO) (err error) {
	
	var user = models.User{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
		ProfileID: request.ProfileID,
	}

	if err = s.userRepo.Save(&user); err != nil {
		return err
	}

	return err
}