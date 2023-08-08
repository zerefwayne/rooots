package config

import (
	"fmt"
	"log"
	"os"

	"github.com/zerefwayne/rooots/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func getDatabaseConnectionURL() string {
	host := os.Getenv("DB_POSTGRESQL_HOST")
	port := os.Getenv("DB_POSTGRESQL_PORT")
	user := os.Getenv("DB_POSTGRESQL_USER")
	password := os.Getenv("DB_POSTGRESQL_PASSWORD")
	dbname := os.Getenv("DB_POSTGRESQL_DATABASE")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func ConnectDB() {
	connectionURL := getDatabaseConnectionURL()

	if db, err := gorm.Open(postgres.Open(connectionURL), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	} else {
		DB = db
		log.Println("database	connection OK")
	}
}

func PingDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("database	ping OK")
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("database	connection close OK")
	}
}

func AutoMigrateModels() {
	if err := DB.AutoMigrate(
		&models.User{},
	); err != nil {
		panic(err)
	}

	log.Println("database	model migration OK")
}
