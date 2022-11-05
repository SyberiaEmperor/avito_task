package main

import (
	"fmt"
	"os"

	"github.com/SyberiaEmperor/avito_task/pkg/handler"
	"github.com/SyberiaEmperor/avito_task/pkg/repository"
	"github.com/SyberiaEmperor/avito_task/pkg/service"
	"github.com/SyberiaEmperor/avito_task/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		return
	}

	if err := godotenv.Load("../.env"); err != nil {
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
		fmt.Printf("Database connection failed\n")
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
