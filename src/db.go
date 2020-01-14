package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getDB() (*gorm.DB, error) {
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPassword := GetEnv("DB_PASSWORD", "none")
	dbPort := GetEnv("DB_PORT", "5432")
	dbUser := GetEnv("DB_USER", "postgres")
	dbName := GetEnv("DB_NAME", "marvin")
	dbSslMode := GetEnv("DB_SSL_MODE", "disable")

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
