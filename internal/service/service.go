package service

import "github.com/NVinox/todos/internal/repository"

type Authorization interface {
}

type Todo interface {
}

type Service struct {
	Authorization
	Todo
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
