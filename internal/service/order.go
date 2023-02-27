package service

import (
	"awesomeProject-L0/internal/model"
	"awesomeProject-L0/internal/repo"
	"log"
)

type OrderService struct {
	repo  repo.Order
	cache map[int]model.Order
}

func NewOrderService(repo repo.Order) *OrderService {

	return &OrderService{repo: repo, cache: make(map[int]model.Order)}
}

func (s *OrderService) CreateOrder(order model.Order) (int, error) {

	// запись данных в БД
	orderId, err := s.repo.CreateOrder(order)
	// запись данных в кэш
	s.cache[orderId] = order

	return orderId, err
}

func (s *OrderService) GetOrderById(orderId int) (model.Order, error) {

	// получение данных из кэша
	order, ok := s.cache[orderId]
	if !ok {
		// получение данных из БД, если данных нет в кэше
		order, err := s.repo.GetOrderById(orderId)
		if err != nil {
			return order, err
		}
		// запись данных в кэш
		s.cache[orderId] = order

		log.Printf("Got order %d from DB\n", orderId)

		return order, err
	}

	log.Printf("Got order %d from CACHE\n", orderId)

	return order, nil
}

func (s *OrderService) UploadCache() error {

	orders, err := s.repo.UploadCache()

	if err != nil {
		return err
	}

	for _, order := range orders {
		s.cache[order.Id] = order
	}
	return nil
}
