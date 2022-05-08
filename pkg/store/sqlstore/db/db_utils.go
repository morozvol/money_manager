package db

import (
	"database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	config "github.com/morozvol/money_manager/pkg/store/sqlstore/config"
	"go.uber.org/zap"
	"time"
)

// New return sqlx.DB
func New(config *config.DBConfig, logger *zap.Logger) (*sqlx.DB, error) {

	// parse connection string
	dbConf, err := pgx.ParseConfig(config.GetConnactionString())
	if err != nil {
		return nil, err
	}

	dbConf.Logger = zapadapter.NewLogger(logger)

	// register pgx conn
	dsn := stdlib.RegisterConnConfig(dbConf)

	sql.Register("wrapper", stdlib.GetDefaultDriver())
	wdb, err := sql.Open("wrapper", dsn)
	if err != nil {
		return nil, err
	}

	db := sqlx.NewDb(wdb, "pgx")
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
