package main

import (
	"github.com/Bruno07/tasks-api/internal/config"
	"github.com/Bruno07/tasks-api/internal/http/routes"
)

func main() {

	config.LoadConfig()

	router := routes.InitRoutes()

	router.Run(":5001")
	
}