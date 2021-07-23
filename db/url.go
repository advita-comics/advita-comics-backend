package db

import (
	"fmt"

	"github.com/advita-comics/advita-comics-backend/config"
)

// MysqlSQLDBDriver - mysql driver name
const MysqlSQLDBDriver = "mysql"

// GenerateMysqlDatabaseURL - generate path to mysql database
func GenerateMysqlDatabaseURL(conf config.DBConfig) string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)
}
