package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

func getDB() (*gorm.DB, error) {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPassword := getEnv("DB_PASSWORD", "none")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbName := getEnv("DB_NAME", "marvin")
	dbSslMode := getEnv("DB_SSL_MODE", "disable")

	configuration := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword,
		dbSslMode,
	)
	db, err := gorm.Open(
		"postgres",
		configuration,
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//DB is either nil or a connection to the configured database using gorm.
var DB, err = getDB()

func init() {
	if err != nil {
		//Don't need to os.exit if DB is not available
		log.Printf("Could not connect to DB\nError: %s", err)
	}
}
