package models

import "time"

type Module struct {
	ID           int64 `gorm:"primaryKey;autoIncrement"`
	ProfileId    int64
	Profile      Profile
	PermissionId int64
	Permission   Permission
	UpdatedAt    time.Time `gorm:"autoUpdateTime:datetime"`
	CreatedAt    time.Time `gorm:"autoCreateTime:datetime"`
}
