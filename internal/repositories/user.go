package repositories

import (
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type UserRepository struct {}

func (ur UserRepository) Save(user *models.User) (*models.User, error) {

	result := db.GetInstance().Model(user).Create(map[string]interface{} {
		"Name": user.Name,
		"Email": user.Email,
		"Password": user.Password,
		"ProfileId": user.ProfileId,
	})

	return user, result.Error

}

