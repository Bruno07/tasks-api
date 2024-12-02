package repositories

import (
	"errors"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type IUserRepository interface {
	Save(request requests.UserRequest) (models.User, error)
	Find(userId int64) (models.User, error)
	GetByEmail(email string) (models.User, error)
}

type UserRepository struct{}

type UserMockRepository struct{}

// Persists task in database
func (ur UserRepository) Save(request requests.UserRequest) (models.User, error) {

	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		ProfileId: request.ProfileId,
	}

	result := db.GetInstance().Create(&user)

	return user, result.Error

}

// Find user by ID
func (ur UserRepository) Find(userId int64) (models.User, error) {

	var user models.User

	db.GetInstance().Where("id = ?", userId).Find(&user)

	return user, nil
}

// Find user by email
func (ur UserRepository) GetByEmail(email string) (models.User, error) {

	var user models.User

	result := db.GetInstance().Where("email = ?", email).Find(&user)

	return user, result.Error
}

// Persists task in database
func (ur UserMockRepository) Save(request requests.UserRequest) (models.User, error) {

	var err error
	user := models.User{}

	if request == (requests.UserRequest{}) {
		err = errors.New("Invalid request!")

	} else {
		user = models.User{
			Name:      request.Name,
			Email:     request.Email,
			Password:  request.Password,
			ProfileId: request.ProfileId,
		}
	}

	return user, err

}

// Find user by ID
func (ur UserMockRepository) Find(userId int64) (models.User, error) {

	userGroup := make(map[int64]models.User)
	userGroup[1] = models.User{ID: 1, Name: "Teste 1", Email: "teste1@email.com"}
	userGroup[2] = models.User{ID: 2, Name: "Teste 2", Email: "teste2@email.com"}

	user := userGroup[userId]

	return user, nil
}

// Find user by email
func (ur UserMockRepository) GetByEmail(email string) (models.User, error) {

	var err error
	var user = models.User{}

	if email == "" {
		err = errors.New("User not found")

	} else {
		user.ID = 1
		user.Name = "master@email.com"
		user.Password = "123456"
	}

	return user, err
}
