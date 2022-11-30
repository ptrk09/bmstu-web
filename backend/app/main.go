package main

import (
	"web/internal/handler"
	"web/internal/repository"
	"web/internal/server"
	"web/internal/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const configPath = "configs/config.yml"

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalln("error init configs:", err.Error())
	}

	db, err := repository.NewPostrgersDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	logrus.Info(viper.GetString("db.port"))
	logrus.Info(viper.GetString("db.host"))

	if err != nil {
		logrus.Fatalln("failed to init db:", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalln("error occured while running http server:", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./app/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
