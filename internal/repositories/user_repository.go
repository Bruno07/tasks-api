package repositories

import (
	"github.com/Bruno07/tasks-api/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Save a user
func (u *UserRepository) Save(user *models.User) error {
	return u.db.Create(&user).Error
}

// Get user by email
func (u UserRepository) GetByEmail(email string) (*models.User, error) {

	var user = models.User{}

	result := u.db.Where("email = ?", email).Find(&user)

	return &user, result.Error

}
