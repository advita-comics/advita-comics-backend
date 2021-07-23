package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/advita-comics/advita-comics-backend/config"
	"github.com/advita-comics/advita-comics-backend/db"
	"github.com/lillilli/vconf"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

var (
	configFile  = flag.String("config", "", "set service config file")
	migratePath = flag.String("migrate-path", "", "path to migrations")
	logLevel    = flag.Int("log-level", 5, "log level")
)

func main() {
	flag.Parse()
	logrus.SetLevel(logrus.Level(*logLevel))

	/* загружаем конфиги */
	cfg := &config.Config{}
	if err := vconf.InitFromFile(*configFile, cfg); err != nil {
		logrus.Fatalf("unable to load config: %s\n", err)
	}

	if cfg.Log.Path != "" {
		logFile, err := os.Create(cfg.Log.Path)
		if err != nil {
			logrus.Fatalf("Unable to open log file: %s", err.Error())
		}
		logrus.SetOutput(logFile)
	}

	conn, err := sql.Open(db.MysqlSQLDBDriver, db.GenerateMysqlDatabaseURL(cfg.DB))
	if err != nil {
		logrus.Fatalf("sql.Open: %s\n", err)
	}
	defer conn.Close()

	if err := goose.SetDialect(db.MysqlSQLDBDriver); err != nil {
		logrus.Fatalf("goose.Up: %s\n", err)
	}
	/* запускаем миграции */
	if err := goose.Up(conn, *migratePath); err != nil {
		logrus.Fatalf("goose.Up: %s\n", err)
	}

	fmt.Println("Migrations complete")
}
