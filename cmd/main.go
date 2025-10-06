package main

import (
	"log"

	"github.com/NVinox/todos"
	"github.com/NVinox/todos/internal/handler"
	"github.com/NVinox/todos/internal/repository"
	"github.com/NVinox/todos/internal/service"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := new(todos.Server)
	if err := server.Run("3000", handler.InitRoutes()); err != nil {
		log.Fatal("Could not to run server: %s", err.Error())
	}
}
