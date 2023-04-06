package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "Admin:Fonteyn@DB(hostname:3306)/slagboom_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Retrieve the license plate from the gate
	licensePlate := "ABC123"

	// Check if the license plate is found in the database
	var name, email string
	err = db.QueryRow("SELECT name, email FROM reservations WHERE license_plate = ? AND start_time <= ? AND end_time >= ?", licensePlate, time.Now(), time.Now()).Scan(&name, &email)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("License plate not found in database.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Welcome %s! Email: %s\n", name, email)
	}
}
