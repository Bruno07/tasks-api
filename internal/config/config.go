package config

import (
	"fmt"
	"os"

	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/infra/queue"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/joho/godotenv"
)

func LoadConfig() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file")
	}

	db.NewDatabase(db.MysqlDatabase{
		DbName:     os.Getenv("DB_NAME"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
	})

	db.GetInstance().AutoMigrate(
		&models.Profile{},
		&models.Permission{},
		&models.User{},
		&models.Module{},
		&models.Task{},
		&models.Notification{},
	)

	fmt.Println(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("QUEUE_USER"),
		os.Getenv("QUEUE_PASSWORD"),
		os.Getenv("QUEUE_HOST"),
		os.Getenv("QUEUE_PORT"),
	))

	queue.Connect(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("QUEUE_USER"),
		os.Getenv("QUEUE_PASSWORD"),
		os.Getenv("QUEUE_HOST"),
		os.Getenv("QUEUE_PORT"),
	))

}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
