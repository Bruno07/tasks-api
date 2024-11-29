package models

import "time"

type Profile struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:25"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:datetime"`
	CreatedAt time.Time `gorm:"autoCreateTime:datetime"`
}