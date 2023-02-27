package handler

import (
	"awesomeProject-L0/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// InitRoutes инициализация эндпойнтов
func (h *Handler) InitRoutes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/order/{id:[0-9]+}", h.getOrderById).Methods("GET")
	router.HandleFunc("/order/", h.startPage)

	return router
}
