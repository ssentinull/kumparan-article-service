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

func ServerPort() string {
	return fmt.Sprintf("%s", os.Getenv("SERVER_PORT"))
}
