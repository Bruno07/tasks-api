package repositories

import (
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type ProfileRepository struct {}

func (pr *ProfileRepository) Find(profile *models.Profile) *models.Profile {

	db.GetInstance().Model(profile).Find(map[string]interface{}{})

	return profile

}
