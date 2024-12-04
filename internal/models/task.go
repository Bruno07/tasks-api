package models

import (
	"errors"
	"time"
)

type Task struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(2500);not null"`
	UserID      int64     `gorm:"type:bigint;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime;type:datetime;not null"`
	UpdateAt    time.Time `gorm:"autoUpdateTime;type:datetime;not null"`
}

func (t *Task) Validate() error {

	if t.Title == "" {
		return errors.New("title is required")
	}

	if t.Description == "" {
		return errors.New("description is required")
	}

	if len(t.Description) > 2500 {
		return errors.New("Number of characters exceeded")
	}

	if t.UserID == 0 {
		return errors.New("User not found")
	}

	return nil

}
