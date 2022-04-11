package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/morozvol/money_manager/pkg/store/sqlstore/config"
	"strings"
	"testing"
	"time"
)

// TestDB ...
var dsn string = ""

func TestDB(t *testing.T, config *config.DBConfig) (*sqlx.DB, func(...string)) {
	t.Helper()
	if dsn == "" {
		if config.Name != "test" {
			t.Fatal("DataBase name != \"test\"")
		}
		// parse connection string
		dbConf, err := pgx.ParseConfig(config.GetConnactionString())
		if err != nil {
			t.Fatal()
		}

		// register pgx conn
		dsn = stdlib.RegisterConnConfig(dbConf)

		sql.Register("wrapper", stdlib.GetDefaultDriver())

	}
	wdb, err := sql.Open("wrapper", dsn)
	if err != nil {
		t.Fatal()
	}

	db := sqlx.NewDb(wdb, "pgx")
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		t.Fatal()
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			_, err := db.Exec(fmt.Sprintf("TRUNCATE %s RESTART IDENTITY CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				t.Errorf("Truncate failed: %s", err.Error())
			}
		}

		err := db.Close()
		if err != nil {
			return
		}
	}
}

func GetTestDBStore(t *testing.T) (*Store, func(...string)) {
	baseConfig, err := config.GetDataBaseConfig("testDb", "../../../config")
	if err != nil {
		t.Fatal("Не удалось получить тестовый конфиг")
	}
	db, truncate := TestDB(t, baseConfig)

	return New(db).(*Store), truncate
}
