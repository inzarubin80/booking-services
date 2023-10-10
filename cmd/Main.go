package main

import (
	"fmt"
	"github.com/inzarubin80/booking-services/pkg/repository"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewMSsqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("pass"),
	})

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	fmt.Println(repos)
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
