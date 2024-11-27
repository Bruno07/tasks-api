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
	UserRepo       repositories.UserRepository
	ModuleRepo     repositories.ModuleRepository
	PermissionRepo repositories.PermissionRepository
	JWTService     JWTService
}

func (us AuthService) Authenticate(request requests.LoginRequest) (*responses.JWTResponse, error) {

	user := us.UserRepo.GetByEmail(request.Email)

	if user == (models.User{}) {
		return nil, errors.New("User not found!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("User not found!")
	}

	response := responses.JWTResponse{
		User: responses.UserResponse{
			Name:  user.Name,
			Email: user.Email,
		},
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		ISS: "tasks-api",
	}

	modules := us.ModuleRepo.GetByProfileId(user.ProfileId)
	for _, module := range modules {
		permission := us.PermissionRepo.Find(module.PermissionId)
		response.User.Permissions = append(response.User.Permissions, permission.Role)
	}

	token, err := us.JWTService.generateToken(response)
	if err != nil {
		return nil, err
	}

	response.AccessToken = *token

	return &response, nil

}
