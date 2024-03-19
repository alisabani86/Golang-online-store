package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_DRIVER := os.Getenv("DRIVER")
	DB_USER := os.Getenv("USER")
	DB_PASSWORD := os.Getenv("PASSWORD")
	DB_HOST := os.Getenv("HOST")
	DB_PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")
	DBDRIVER := os.Getenv("DBDRIVER")
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", DB_DRIVER, DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DBNAME)

	db, err := sql.Open(DBDRIVER, connectionString)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil

}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
