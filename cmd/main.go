package main

import (
	"log"

	"github.com/NVinox/todos"
)

func main() {
	server := new(todos.Server)

	if err := server.Run("3000"); err != nil {
		log.Fatal("Could not to run server: %s", err.Error())
	}
}
