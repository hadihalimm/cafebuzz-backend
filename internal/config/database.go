package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Close() error
	Create(value interface{}) error
}

type database struct {
	db *gorm.DB
}

var (
	user       = os.Getenv("DB_USER")
	password   = os.Getenv("DB_PASSWORD")
	host       = os.Getenv("DB_HOST")
	port       = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
	dbInstance *database
)

func ConnectToDatabase() Database {
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, password, host, port, dbName)
	gormDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to database")
	dbInstance := &database{
		db: gormDB,
	}
	return dbInstance
}

func (d *database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	return sqlDB.Close()
}

func (d *database) Create(value interface{}) error {
	return d.db.Create(value).Error
}
