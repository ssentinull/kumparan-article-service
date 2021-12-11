package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/kumparan-article-service/config"
)

func main() {
	mode := flag.String("mode", "", "")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s/db/migration", cwd),
		config.PostgresDSN())

	if err != nil {
		logrus.Fatal(err)
	}

	if *mode == "up" {
		err = m.Up()
	} else {
		err = m.Down()
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
