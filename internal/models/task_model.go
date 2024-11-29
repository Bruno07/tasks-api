package models

import "time"

type Task struct {
	ID            int64      `gorm:"primaryKey;autoIncrement"`
	Summary       string     `gorm:"type:varchar(2500);not null"`
	PerformedDate *time.Time `gorm:"autoCreateTime;type:datetime;not null"`
	UserID        int64
	User          User
}
