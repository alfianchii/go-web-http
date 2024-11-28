package config

import (
	"fmt"
)

const (
	DBHost = "localhost"
	DBPort = 5432
	DBUser = "your-username"
	DBPassword = "your-password"
	DBName = "your-db-name"
)

func PSQLInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPassword, DBName)
}