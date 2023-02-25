package handler

import (
	"awesomeProject-L0/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// InitRoutes инициализация эндпойнтов
func (h *Handler) InitRoutes() *gin.Engine { //переписать на net/http
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/:id", h.getOrderById)
		//api.POST("/", h.createOrder)
	}
	return router
}
