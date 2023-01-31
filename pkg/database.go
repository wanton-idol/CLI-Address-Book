package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	en "github.com/wanton-idol/cli-address-book/entity"
)

func DBConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &en.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	return db

}

func NewConnection(config *en.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	return db, nil
}
