package repositories

import (
	"fmt"

	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type IModuleRepository interface {
	GetByProfileId(profileId int64) []models.Module
}

type ModuleRepository struct{}

type ModuleMockRepository struct{}

func (mr *ModuleRepository) GetByProfileId(profileId int64) []models.Module {

	var modules []models.Module

	db.GetInstance().Where("profile_id = ?", profileId).Find(&modules)

	return modules

}

func (mr *ModuleMockRepository) GetByProfileId(profileId int64) []models.Module {

	var modules []models.Module

	var groupPermissionsProfile = map[int64][]int64{
		1: []int64{1,2,3},
		2: []int64{1,2,3,4},
	}

	for _, permission := range groupPermissionsProfile[profileId] {

		module := models.Module{
			ProfileId: profileId,
			PermissionId: permission,
		}

		modules = append(modules, module)
	}

	fmt.Println(modules)

	return modules

}
