package models

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(140);not null"`
	Email     string    `gorm:"type:varchar(200);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	ProfileID int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;type:datetime;not null"`
	UpdateAt  time.Time `gorm:"autoUpdateTime;type:datetime;not null"`
}

// Validate user model
func (u *User) Validate() error {

	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}

	if !isValidEmail(u.Email) {
		return errors.New("Invalid email format")
	}

	if len(u.Password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	return nil

}

// Check email validity
func isValidEmail(email string) bool {
	regex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)

	return re.MatchString(email)
}

// Creates a hash of the password
func (u *User) HashPassword() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil

}

// Check password
func (u *User) CheckPassword(password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil

}
