package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

var (
	// user     = os.Getenv("DB_USER")
	// password = os.Getenv("DB_PASSWORD")
	// host     = os.Getenv("DB_HOST")
	// port     = os.Getenv("DB_PORT")
	// dbName   = os.Getenv("DB_NAME")

	userLocal     = os.Getenv("DB_USER_LOCAL")
	passwordLocal = os.Getenv("DB_PASSWORD_LOCAL")
	hostLocal     = os.Getenv("DB_HOST_LOCAL")
	portLocal     = os.Getenv("DB_PORT_LOCAL")
	dbNameLocal   = os.Getenv("DB_NAME_LOCAL")
	dbInstance    *Database
)

func ConnectToDatabase() *Database {
	if dbInstance != nil {
		return dbInstance
	}
	// connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, password, host, port, dbName)
	connStrLocal := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", userLocal, passwordLocal, hostLocal, portLocal, dbNameLocal)
	gormDB, err := gorm.Open(postgres.Open(connStrLocal), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to database")
	dbInstance := &Database{
		Gorm: gormDB,
	}
	return dbInstance
}

func (d *Database) Close() error {
	sqlDB, err := d.Gorm.DB()
	if err != nil {
		log.Fatal(err)
	}
	return sqlDB.Close()
}

func (d *Database) AutoMigrate(value ...interface{}) error {
	return d.Gorm.AutoMigrate(value...)
}

func (d *Database) DropTable(dst ...interface{}) error {
	return d.Gorm.Migrator().DropTable(dst...)
}
