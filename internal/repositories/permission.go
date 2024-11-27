package repositories

import (
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type PermissionRepository struct {}

func (pr *PermissionRepository) Find(permissionId int64) models.Permission {

	var permission models.Permission

	db.GetInstance().Where("id = ?", permissionId).Find(&permission)

	return permission

}
