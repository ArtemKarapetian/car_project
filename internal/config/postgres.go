package config

import (
	"fmt"
	"os"
)

func GeneratePostgresDSN() (string, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || dbname == "" {
		return "", fmt.Errorf("missing required env vars")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname), nil
}
