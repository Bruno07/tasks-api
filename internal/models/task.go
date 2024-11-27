package models

import "time"

type Task struct {
	Summary       string `gorm:"type:varchar(2500);not null"`
	PerformedDate *time.Time
	UserID        uint
	User          User
}
