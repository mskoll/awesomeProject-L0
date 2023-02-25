package service

import (
	"awesomeProject-L0/internal/model"
	"awesomeProject-L0/internal/repo"
)

type OrderService struct {
	repo repo.Order
}

func NewOrderService(repo repo.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order model.Order) (int, error) {
	return s.repo.CreateOrder(order)
}

func (s *OrderService) GetOrderById(orderId int) (model.Order, error) {
	return s.repo.GetOrderById(orderId)
}
