package services

import (
	"errors"
	"time"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/http/responses"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo       repositories.UserRepository
	moduleRepo     repositories.ModuleRepository
	permissionRepo repositories.PermissionRepository
	jwtService     JWTService
}

func (us AuthService) Authenticate(request requests.LoginRequest) (responses.JWTResponse, error) {

	user, err := us.userRepo.GetByEmail(request.Email)
	if err != nil {
		return responses.JWTResponse{}, err
	}

	if user == (models.User{}) {
		return responses.JWTResponse{}, errors.New("User not found!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return responses.JWTResponse{}, errors.New("User not found!")
	}

	response := responses.JWTResponse{
		User: responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		ISS:       "tasks-api",
	}

	modules := us.moduleRepo.GetByProfileId(user.ProfileId)
	for _, module := range modules {
		permission, err := us.permissionRepo.Find(module.PermissionId)
		if err != nil {
			return responses.JWTResponse{}, err
		}

		response.User.Permissions = append(response.User.Permissions, permission.Role)
	}

	token, err := us.jwtService.generateToken(response)
	if err != nil {
		return responses.JWTResponse{}, err
	}

	response.AccessToken = *token

	return response, nil

}
