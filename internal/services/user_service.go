package services

import (
	"errors"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/http/responses"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
)

type UserService struct {
	userRepo       repositories.IUserRepository
	profileRepo    repositories.ProfileMockRepository
	permissionRepo repositories.PermissionMockRepository
	moduleRepo     repositories.ModuleMockRepository
}

func NewUserService(
	userRepo repositories.IUserRepository,
	profileRepo repositories.ProfileMockRepository,
	permissionRepo repositories.PermissionMockRepository,
	moduleRepo repositories.ModuleMockRepository,
) *UserService {
	return &UserService{
		userRepo:       userRepo,
		profileRepo:    profileRepo,
		permissionRepo: permissionRepo,
		moduleRepo:     moduleRepo,
	}
}

func (us UserService) Create(request requests.UserRequest) (responses.UserResponse, error) {

	profile, err := us.profileRepo.Find(request.ProfileId)
	if err != nil {
		return responses.UserResponse{}, err
	}

	if profile == (models.Profile{}) {
		return responses.UserResponse{}, errors.New("Profile not found!")
	}

	userCreated, err := us.userRepo.Save(request)
	if err != nil {
		return responses.UserResponse{}, err
	}

	response := responses.UserResponse{
		Name:  userCreated.Name,
		Email: userCreated.Email,
	}

	// modules := us.moduleRepo.GetByProfileId(profile.ID)
	// for _, module := range modules {
	// 	permission, err := us.permissionRepo.Find(module.PermissionId)
	// 	if err != nil {
	// 		return responses.UserResponse{}, err
	// 	}
		
	// 	response.Permissions = append(response.Permissions, permission.Role)
	// }

	return response, err

}
