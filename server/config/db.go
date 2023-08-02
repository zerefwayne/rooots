package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/zerefwayne/rooots/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Err error

func SetupDB() {

	host := os.Getenv("DB_POSTGRESQL_HOST")
	port := os.Getenv("DB_POSTGRESQL_PORT")
	user := os.Getenv("DB_POSTGRESQL_USER")
	password := os.Getenv("DB_POSTGRESQL_PASSWORD")
	dbname := os.Getenv("DB_POSTGRESQL_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Printf("Connecting to DB on: %s\n", dsn)

	DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic(Err)
	}

	log.Println("Successfully connected to database!")
}

func AutoMigrateDB() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	log.Println("Successfully migrated all tables to database!")
}
