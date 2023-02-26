package main

import (
	"awesomeProject-L0/internal"
	"awesomeProject-L0/internal/handler"
	"awesomeProject-L0/internal/nats"
	"awesomeProject-L0/internal/repo"
	"awesomeProject-L0/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {

	// инициализация конфига
	if err := initConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	// подключение к БД
	// передаем данные для подключения
	db, err := repo.Init(repo.Conf{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatalf("DB-init error: %s", err.Error())
	}

	// инициализация репозитория для работы с БД
	repos := repo.NewRepo(db)
	// инициализация сервиса - бизнес-логика
	services := service.NewService(repos)
	// инициализация хэндлера
	handlers := handler.NewHandler(services)

	// загрузка кэша из БД
	err = services.UploadCache()
	if err != nil {
		log.Fatalf("Cache error: %s", err.Error())
	}

	// подключение к nats-streaming
	stanConn, err := nats.Init(nats.Conf{
		Cluster: viper.GetString("stan.cluster"),
		Client:  viper.GetString("stan.client"),
	})
	if err != nil {
		log.Fatalf("STAN error %s", err.Error())
	}

	// инициализация подписчика nats-streaming
	sub := nats.NewSubscriber(stanConn, services)
	// подписка на канал
	nats.Subscribe(sub)
	// инициализация издателя nats-streaming
	pub := nats.NewPublisher(stanConn)
	// публикация в канал
	nats.Publish(pub)

	// инициализация сервера
	server := new(internal.Server)
	// запуск сервера
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}

}

// initConfig иницаиализация конфига
func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
