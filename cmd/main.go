package main

import (
	"log"

	"github.com/NVinox/todos"
	"github.com/NVinox/todos/internal/handler"
)

func main() {
	handler := new(handler.Handler)

	server := new(todos.Server)
	if err := server.Run("3000", handler.InitRoutes()); err != nil {
		log.Fatal("Could not to run server: %s", err.Error())
	}
}
