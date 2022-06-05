//go:build (darwin && cgo) || linux

package lib

import (
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const postgresqlDriver = "postgres"

func NewPostgresqlConnection(cfg DatabaseConfig) (*sqlx.DB, error) {
	dbConnection, err := sqlx.Open(postgresqlDriver, cfg.DSN)
	if err != nil {
		return nil, err
	}
	dbConnection.SetConnMaxIdleTime(cfg.MaxIdleDuration)
	dbConnection.SetMaxOpenConns(cfg.MaxOpenConnections)
	dbConnection.SetConnMaxLifetime(cfg.MaxLifeTimeDuration)
	dbConnection.SetMaxIdleConns(cfg.MaxIdleConnections)

	return dbConnection, nil
}
