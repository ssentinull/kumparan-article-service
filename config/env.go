package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	env := os.Getenv("ENV")
	if env != "dev" && env != "development" && env != "" {
		logrus.Warn("running using OS env variables")

		return
	}

	if err := godotenv.Load(); err != nil {
		logrus.Warn(".env file not found")

		return
	}

	logrus.Warn("running using .env file")

	return
}

func Env() string {
	return fmt.Sprintf("%s", os.Getenv("ENV"))
}

func PostgresDSN() string {
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		username,
		password,
		host,
		port,
		dbName)
}

func ServerPort() string {
	return fmt.Sprintf("%s", os.Getenv("SERVER_PORT"))
}
