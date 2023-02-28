package main

import (
	"awesomeProject-L0/internal/nats"
	"github.com/spf13/viper"
	"log"
)

func main() {

	// инициализация конфига
	if err := initConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	// подключение к nats-streaming
	stanConn, err := nats.Init(nats.Conf{
		Cluster: viper.GetString("stan.cluster"),
		Client:  viper.GetString("stan.clientP"),
	})
	if err != nil {
		log.Fatalf("STAN error %s", err.Error())
	}

	// инициализация издателя nats-streaming
	pub := NewPublisher(stanConn)
	// публикация в канал
	Publish(pub)
}

// initConfig иницаиализация конфига
func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
