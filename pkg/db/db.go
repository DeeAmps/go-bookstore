package db

import (
	"log"

	"github.com/DeeAmps/go-bookstore-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	config := config.GetConfigurations()
	dsn := "host=" + config.DBHost + " user=" + config.DBUser + " dbname=" + config.DBName + " password=" + config.DBPass + " sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
