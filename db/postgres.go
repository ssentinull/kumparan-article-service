package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/kumparan-article-service/config"
)

func NewDBConn() *sql.DB {
	db, err := sql.Open("postgres", config.PostgresDSN())
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}
