package models

import "time"

type Notification struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Payload   string `gorm:"type:longtext"`
	UserID    int64
	CreatedAt time.Time `gorm:"autoCreateTime;datetime"`
	UpdateAt  time.Time `gorm:"autoUpdateTime;datetime"`
}
