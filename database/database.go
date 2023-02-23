package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	// dbURL := "postgresql://user:password@postgres:5432/crud"

	dbURL := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Uncomment the following lines of code, if you want to run migrations

	// log.Println("Running Migrations")
	// TODO: Add migrations
	// db.AutoMigrate(&models.CustomerCompany{}, &models.Customer{}, &models.Order{}, &models.OrderItem{}, &models.Delivery{})

	Database = DBInstance{DB: db}
}
