package main

import (
	"github.com/Bruno07/tasks-api/internal/config"
	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
	"github.com/Bruno07/tasks-api/internal/services"
)

func SeederUser() {

	config.LoadConfig()

	var requests = []requests.UserRequestDTO{
		{Name: "Master", Email: "master@email.com", Password: "12345678", ProfileID: 1},
		{Name: "Tec 1", Email: "tec1@email.com", Password: "12345678", ProfileID: 2},
		{Name: "Tec 2", Email: "tec2@email.com", Password: "12345678", ProfileID: 2},
	}

	repository := repositories.NewUserRepository(db.GetInstance())
	service := services.NewUserService(repository)

	for _, request := range requests {
		service.Create(&request)
	}

}

func main() {
	SeederUser()
}
