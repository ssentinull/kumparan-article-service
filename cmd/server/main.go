package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ssentinull/kumparan-article-service/config"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	s := &http.Server{
		Addr:         ":" + config.ServerPort(),
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	log.Fatal(e.StartServer(s))
}
