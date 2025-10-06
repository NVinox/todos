package main

import (
	"log"

	"github.com/NVinox/todos"
	"github.com/NVinox/todos/internal/handler"
	"github.com/NVinox/todos/internal/repository"
	"github.com/NVinox/todos/internal/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Could not read config file: %s", err.Error())
	}

	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := new(todos.Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal("Could not to run server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
