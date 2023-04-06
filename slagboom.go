package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() error {
	//ToDo: make configurable
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "my-secret-pw",
		Net:       "tcp",
		Addr:      "127.0.0.1:3308",
		DBName:    "fonteyn_reservations",
		ParseTime: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	return nil
}

type ValidationResult struct {
	Result bool
	Message string
}

type period struct {
	startDate time.Time
	endDate   time.Time
}
