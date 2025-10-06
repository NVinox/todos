package handler

import (
	"github.com/NVinox/todos/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	api.POST("/todos", h.createTodo)
	api.GET("/todos", h.getTodos)
	api.GET("/todos/:id", h.getTodo)
	api.PATCH("/todos/:id", h.updateTodo)
	api.DELETE("/todos/:id", h.deleteTodo)

	auth := router.Group("/auth")
	auth.POST("/sign-up", h.signUp)
	auth.POST("/sign-in", h.signIn)

	return router
}
