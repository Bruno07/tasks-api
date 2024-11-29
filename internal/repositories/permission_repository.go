package repositories

import (
	"errors"

	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type IPermissionRepository interface {
	Find(permissionId int64) (models.Permission, error)
}

type PermissionRepository struct {}

type PermissionMockRepository struct {}

func (pr *PermissionRepository) Find(permissionId int64) (models.Permission, error) {

	var permission models.Permission

	result := db.GetInstance().Where("id = ?", permissionId).Find(&permission)

	return permission, result.Error

}

func (pr *PermissionMockRepository) Find(permissionId int64) (models.Permission, error) {

	var err error
	var permission = models.Permission{}

	switch permissionId {
		case 1:
			permission.ID = 1
			permission.Role = "CREATE"
		case 2:
			permission.ID = 2
			permission.Role = "UPDATE"
		case 3:
			permission.ID = 3
			permission.Role = "VIEW"
		case 4:
			permission.ID = 4
			permission.Role = "DELETE"
		default:
			err = errors.New("Invalid permission!")
	}

	return permission, err

}
