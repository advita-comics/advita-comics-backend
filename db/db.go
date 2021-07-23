package db

import (
	"context"
	"time"

	"github.com/advita-comics/advita-comics-backend/db/dao"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// DB - database interface
type DB interface {
	Dao() dao.Models
	Repo() rel.Repository
	Close()
}

type db struct {
	repository rel.Repository
	adapter    *mysql.Adapter
	models     dao.Models
}

// New - return new db instance
func New(connectionURL string, loglevel int64) DB {
	adapter, err := mysql.Open(connectionURL)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	repo := rel.New(adapter)
	repo.Instrumentation(func(ctx context.Context, op string, message string) func(err error) {
		t := time.Now()

		if logrus.Level(loglevel) == logrus.DebugLevel {
			duration := time.Since(t)
			logrus.Debug("[duration: ", duration, " op: ", op, "] ", message, " err: ", err)
		}

		return func(err error) {
			duration := time.Since(t)
			if err != nil {
				logrus.Error("[duration: ", duration, " op: ", op, "] ", message, " err:", err)
			}
		}
	})

	return &db{repository: repo, adapter: adapter, models: dao.NewModels(repo)}
}

// Repo - возвращает инстанс рела
func (db *db) Repo() rel.Repository {
	return db.repository
}

// Close - close connection to db
func (db *db) Close() {
	db.adapter.Close()
}

func (db *db) Dao() dao.Models {
	return db.models
}
