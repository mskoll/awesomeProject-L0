package service

import (
	"awesomeProject-L0/internal/model"
	"awesomeProject-L0/internal/repo"
)

type Order interface {
	CreateOrder(order model.Order) (int, error)
	GetOrderById(orderId int) (model.Order, error)
}
type Service struct {
	Order
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Order: NewOrderService(repo.Order),
	}
}
