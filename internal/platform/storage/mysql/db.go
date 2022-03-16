package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manuhdez/golang-api-hex/internal/platform/config"
)

func Connect(dbConfig config.DbConfig) (*sql.DB, error) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	return sql.Open("mysql", dbUri)
}
