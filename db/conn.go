package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(host string, port string, user string, password string, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get generic DB interface:", err)
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println("Failed to ping database:", err)
		return nil, err
	}

	fmt.Println("Connected to " + dbname)
	return db, nil

}