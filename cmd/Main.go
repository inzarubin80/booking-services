package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/handler"
	"github.com/inzarubin80/booking-services/pkg/repository"
	"github.com/inzarubin80/booking-services/pkg/service"
	"github.com/spf13/viper"
)

// @title Booking Services API
// @version 1.0
// @description API Server for Booking Services

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewMSsqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.pass"),
		Instance: viper.GetString("db.instance"),
		Port: viper.GetString("db.port"),	
	})

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(booking.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("booking-services started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
