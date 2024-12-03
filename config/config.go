package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	Port = 3333
	Host = "localhost"
	DBHost = "127.0.0.1"
	DBPort = 5432
	DBUser = "your-username"
	DBName = "your-database"
)

var Address = fmt.Sprintf("%s:%d", Host, Port)

func dbConfig () string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBName)
}

func DBConnect() *sqlx.DB {
	db, err := sqlx.Open("postgres", dbConfig())
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