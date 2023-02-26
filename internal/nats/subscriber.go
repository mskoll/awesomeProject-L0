package nats

import (
	"awesomeProject-L0/internal/model"
	"awesomeProject-L0/internal/service"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

type Subscriber struct {
	connection *stan.Conn
	service    *service.Service
}

func NewSubscriber(conn *stan.Conn, serv *service.Service) *Subscriber {
	return &Subscriber{
		connection: conn,
		service:    serv,
	}
}

// Subscribe подписка на канал
func Subscribe(sub *Subscriber) {
	(*sub.connection).Subscribe("order", func(msg *stan.Msg) {

		var order model.Order

		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Not valid data error: %s", err.Error())
		} else {
			// отправка полученного order в сервис
			_, err := sub.service.CreateOrder(order)
			if err != nil {
				log.Printf("Create order error: %s", err.Error())
			}
		}

	})
}
