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

type MockUserRepository struct {
	MockSave func (user *models.User) error
	MockGetByEmail func (email string) (*models.User, error)
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

// Save a user (Mock)
func (m MockUserRepository) Save(user *models.User) error {
	return m.MockSave(user)
}

// Get user by email (Mock)
func (m MockUserRepository) GetByEmail(email string) (*models.User, error) {
	return m.MockGetByEmail(email)
}
