package repository

import (
	"fmt"

	"github.com/SyberiaEmperor/avito_task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	conn, cerr := db.DB()

	if err = conn.Ping(); err != nil || cerr != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transaction{})

	return db, err
}
