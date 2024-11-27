package repositories

import (
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type ModuleRepository struct {}

func (mr *ModuleRepository) GetByProfileId(profileId int64) []models.Module {

	var modules []models.Module

	db.GetInstance().Where("profile_id = ?", profileId).Find(&modules)

	return modules

}
