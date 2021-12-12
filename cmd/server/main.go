package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/kumparan-article-service/config"
	"github.com/ssentinull/kumparan-article-service/db"
	_articleHndlr "github.com/ssentinull/kumparan-article-service/pkg/article/handler/http"
	_articleRepo "github.com/ssentinull/kumparan-article-service/pkg/article/repository/postgres"
	_articleUcase "github.com/ssentinull/kumparan-article-service/pkg/article/usecase"
	"github.com/ssentinull/kumparan-article-service/pkg/utils/big_cache"
)

func initLogger() {
	logLevel := logrus.ErrorLevel

	switch config.Env() {
	case "dev", "development":
		logLevel = logrus.InfoLevel
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		DisableSorting:  true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05 02-01-2006",
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetLevel(logLevel)
}

func init() {
	initLogger()
}

func main() {
	e := echo.New()
	db := db.NewDBConn()
	bigCache := big_cache.NewBigCache(big_cache.Config{EvictionTime: time.Duration(5)})
	cacher := big_cache.NewCacher(&big_cache.CacherConfig{BigCache: bigCache})

	articleRepo := _articleRepo.NewArticleRepository(db, cacher)
	articleUsecase := _articleUcase.NewArticleUsecase(articleRepo)
	_articleHndlr.NewArticleHandler(e, articleUsecase)

	s := &http.Server{
		Addr:         ":" + config.ServerPort(),
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	log.Fatal(e.StartServer(s))
}
