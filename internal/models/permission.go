package models

import "time"

type Permission struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Role      string    `gorm:"size:25"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:datetime"`
	CreatedAt time.Time `gorm:"autoCreateTime:datetime"`
}
