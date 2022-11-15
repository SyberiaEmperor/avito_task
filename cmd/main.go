package main

import (
	"os"

	"github.com/SyberiaEmperor/avito_task/pkg/handler"
	"github.com/SyberiaEmperor/avito_task/pkg/repository"
	"github.com/SyberiaEmperor/avito_task/pkg/service"
	"github.com/SyberiaEmperor/avito_task/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s",err.Error())
		return
	}

	if err := godotenv.Load("../.env"); err != nil {
		logrus.Fatalf("error initializing environment: %s",err.Error())
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),

		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Database connection failed: %s",err.Error())
		return
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	srv := new(server.Server)

	if err := srv.Run(viper.GetString("server.port"), handler.InitRoutes()); err != nil {
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
