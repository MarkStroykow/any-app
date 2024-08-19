package main

import (
	"fmt"
	"log"

	"github.com/MarkStroykow/any-app"
	"github.com/MarkStroykow/any-app/pkg/handler"
	"github.com/MarkStroykow/any-app/pkg/repository"
	"github.com/MarkStroykow/any-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initConfig: %s", err)
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	test := viper.GetString("port")
	fmt.Println(test)
	serv := new(any.Server)
	if err := serv.Run(handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
