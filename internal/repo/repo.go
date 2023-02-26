package repo

import (
	"awesomeProject-L0/internal/model"
	"github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(order model.Order) (int, error)
	GetOrderById(orderId int) (model.Order, error)
	UploadCache() ([]model.Order, error)
}
type Repo struct {
	Order
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Order: NewOrderDB(db),
	}
}
