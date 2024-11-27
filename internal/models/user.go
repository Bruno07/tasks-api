package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:150;not null"`
	Email     string `gorm:"size:200;not null"`
	Password  string `gorm:"size"`
	ProfileId int64
	Profile   Profile   
	UpdatedAt time.Time `gorm:"autoUpdateTime:datetime"`
	CreatedAt time.Time `gorm:"autoCreateTime:datetime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if len(u.Password) > 0 && !isHashed(u.Password) {

		password, err := generatePassword(u.Password)

		if err != nil {
			return err
		}

		u.Password = password
	}

	return nil

}

func isHashed(password string) bool {
	return len(password) > 0 && (password[:4] == "$2a$" || password[:4] == "$2b$" || password[:4] == "$2y$")
}

func generatePassword(password string) (string, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashPassword), err

}
