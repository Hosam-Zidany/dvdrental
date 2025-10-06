package database

import (
	"fmt"
	"log"

	"github.com/Hosam-Zidany/dvdrental/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InintDB(cfg config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // disable plural table names
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	DB = db

	log.Println("database connected")
}
