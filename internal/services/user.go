package services

import (
	"errors"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/http/responses"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
)

type UserService struct {
	UserRepo       repositories.UserRepository
	ProfileRepo    repositories.ProfileRepository
	PermissionRepo repositories.PermissionRepository
	ModuleRepo     repositories.ModuleRepository
}

func NewUserService() UserService {
	return UserService{
		UserRepo:       repositories.UserRepository{},
		ProfileRepo:    repositories.ProfileRepository{},
		PermissionRepo: repositories.PermissionRepository{},
		ModuleRepo:     repositories.ModuleRepository{},
	}
}

func (us UserService) Create(request requests.UserRequest) (*responses.UserResponse, error) {

	profile := us.ProfileRepo.Find(&models.Profile{ID: request.ProfileId})

	if profile == (&models.Profile{}) {
		return nil, errors.New("Profile not found!")
	}

	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		ProfileId: profile.ID,
	}

	userCreated, err := us.UserRepo.Save(&user)
	if err != nil {
		return nil, err
	}

	response := responses.UserResponse{
		Name: userCreated.Name,
		Email: userCreated.Email,
	}

	modules := us.ModuleRepo.GetByProfileId(profile.ID)
	for _, module := range modules {
		permission := us.PermissionRepo.Find(module.PermissionId)
		response.Permissions = append(response.Permissions, permission.Role)
	}
	
	return &response, err

}
