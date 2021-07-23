package main

import (
	"flag"
	"os"

	"github.com/advita-comics/advita-comics-backend/config"
	"github.com/advita-comics/advita-comics-backend/db"
	"github.com/advita-comics/advita-comics-backend/http"
	"github.com/lillilli/vconf"
	"github.com/sirupsen/logrus"
)

var (
	configFile = flag.String("config", "", "set service config file")
)

func main() {
	flag.Parse()

	/* загружаем конфиги */
	cfg := &config.Config{}
	if err := vconf.InitFromFile(*configFile, cfg); err != nil {
		logrus.Fatalf("unable to load config: %s\n", err)
	}

	logrus.SetLevel(logrus.Level(cfg.Log.Level))

	if cfg.Log.Path != "" {
		logFile, err := os.Create(cfg.Log.Path)
		if err != nil {
			logrus.Fatalf("Unable to open log file: %s", err.Error())
		}
		logrus.SetOutput(logFile)
	}

	/* подключаемся к базе */
	dbConn := db.New(db.GenerateMysqlDatabaseURL(cfg.DB), cfg.Log.Level)
	defer dbConn.Close()

	/* стартуем сервер */
	http.NewServer(cfg, dbConn).Start()
}
