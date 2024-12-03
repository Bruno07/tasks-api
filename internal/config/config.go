package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {

	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}