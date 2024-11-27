package repositories

import (
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type UserRepository struct{}

func (ur UserRepository) Save(user *models.User) (*models.User, error) {

	result := db.GetInstance().Create(&user)

	return user, result.Error

}

func (ur UserRepository) GetByEmail(email string) models.User {

	var user models.User

	db.GetInstance().Where("email = ?", email).Find(&user)

	return user
}
