package main

import (
	"awesomeProject-L0/internal"
	"awesomeProject-L0/internal/handler"
	"awesomeProject-L0/internal/nats"
	"awesomeProject-L0/internal/repo"
	"awesomeProject-L0/internal/service"
	"context"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"time"
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
	log.Printf("DB connected\n")

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
	log.Printf("Cache uploaded\n")

	// подключение к nats-streaming
	stanConn, err := nats.Init(nats.Conf{
		Cluster: viper.GetString("stan.cluster"),
		Client:  viper.GetString("stan.clientS"),
	})
	if err != nil {
		log.Fatalf("STAN error %s", err.Error())
	}
	log.Printf("STAN connected\n")

	// инициализация подписчика nats-streaming
	sub := nats.NewSubscriber(stanConn, services)
	// подписка на канал
	nats.Subscribe(sub)

	// инициализация сервера
	server := new(internal.Server)

	// запуск сервера
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Printf("Server error: %s", err.Error())
		}
	}()
	log.Printf("Server started\n")

	// канал для получения сигналов системы
	stop := make(chan os.Signal, 1)
	// получение сигнала, что приложение завершилось
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("DB connection close error: %s", err.Error())
	}

	if err := nats.Close(*stanConn); err != nil {
		log.Printf("STAN connection close error: %s", err.Error())
	}

}

// initConfig иницаиализация конфига
func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
