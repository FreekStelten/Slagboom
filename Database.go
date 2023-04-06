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
		User:      "Admin",
		Passwd:    "Fonteyn@DB",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "slagboom_db",
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
	Result  bool
	Message string
}

type period struct {
	startDate time.Time
	endDate   time.Time
}

func CheckBooking(licensePlate string, time time.Time, countryCode string) (ValidationResult, error) {
	result := ValidationResult{Result: false, Message: ""}
	rows, err := db.Query("SELECT StartDate, EndDate FROM Booking B INNER JOIN Customer C ON C.ID = B.CustomerID INNER JOIN Parc P ON P.ID = B.ParcID WHERE C.LicensePlate = ? AND P.Country = ?", licensePlate, countryCode)
	if err != nil {
		return result, fmt.Errorf("unable to execute query %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		var period period
		if err := rows.Scan(&period.startDate, &period.endDate); err != nil {
			return result, fmt.Errorf("failed to read data from rows %v", err)
		}
		if period.startDate.Before(time) && period.endDate.After(time) {
			result.Result = true
			return result, nil
		}
	}

	result.Result = false
	result.Message = fmt.Sprintf("no active booking found for %v", licensePlate)
	return result, nil
}
