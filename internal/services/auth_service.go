package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/Bruno07/tasks-api/internal/http/auth"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
)

type AuthService struct {
	userRepo repositories.IUserRepository	
}

func NewAuthService(userRepo repositories.IUserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(email string, password string) (string, int64, error) {

	fmt.Println("Aqui 1")

	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", 0, err
	}

	fmt.Println("aqui 2")

	if user == (&models.User{}) {
		return "", 0, errors.New("User not found!")
	}

	if !user.CheckPassword(password){
		return "", 0, errors.New("User not found!")
	}

	var JWT = auth.JWT{}
	var expiresAt = time.Now().Add(1 * time.Hour).Unix()
	token, err := JWT.GenerateToken(user, expiresAt, "tasks-api", user.GetPermissions())
	if err != nil {
		return "", 0, err
	}

	return token, expiresAt, nil
}