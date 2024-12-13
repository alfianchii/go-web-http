package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Address = fmt.Sprintf("%s:%s", GetENV("APP_URL"), GetENV("APP_PORT"))

func InitENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetENV(key string) string {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error reading .env file")
	}

	return myEnv[key]
}

func dbConfig () string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", GetENV("DB_HOST"), GetENV("DB_PORT"), GetENV("DB_USERNAME"), GetENV("DB_DATABASE"))
}

func DBConnect() *sqlx.DB {
	db, err := sqlx.Open(GetENV("DB_CONNECTION"), dbConfig())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v\n", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")
	return db
}