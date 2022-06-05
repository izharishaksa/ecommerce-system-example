//go:build (darwin && cgo) || linux

package lib

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const mysqlDriver = "mysql"

func NewMySqlConnection(cfg DatabaseConfig) (*sqlx.DB, error) {
	dbConnection, err := sqlx.Open(mysqlDriver, cfg.DSN)
	if err != nil {
		return nil, err
	}
	dbConnection.SetConnMaxIdleTime(cfg.MaxIdleDuration)
	dbConnection.SetMaxOpenConns(cfg.MaxOpenConnections)
	dbConnection.SetConnMaxLifetime(cfg.MaxLifeTimeDuration)
	dbConnection.SetMaxIdleConns(cfg.MaxIdleConnections)
	return dbConnection, nil
}
