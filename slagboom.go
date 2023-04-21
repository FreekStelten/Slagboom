package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set up database connection parameters
	dbUser := "Admin"
	dbPass := "Fonteyn@DB"
	dbName := "slagboom_db"
	dbAddress := "127.0.0.1"

	plate := flag.String("plate", "", "er moet een kenteken opgegeven worden!")
	flag.Parse()
	if !flag.Parsed() || *plate == "" {
		flag.Usage()
		log.Println("Geen kenteken opgegeven, probeer het opnieuw.")
		logError("Geen kenteken opgegeven, probeer het opnieuw.")
		os.Exit(1)
	}

	// Create data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbAddress, dbName)

	// Connect to database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		errMsg := fmt.Sprintf("er kan geen connecting naar de database gemaakt worden: %s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}
	defer db.Close()

	// Ping database to ensure connection is valid
	err = db.Ping()
	if err != nil {
		errMsg := fmt.Sprintf("Er kan niet gepinged worden naar de database: %s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}

	// Connection successful
	fmt.Println("Connected to database!")
	var name, licenseplate string
	rows, err := db.Query("SELECT name,licenseplate FROM klant WHERE licenseplate = ?", *plate)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&name, &licenseplate)
		if err != nil {
			errMsg := fmt.Sprintf("%s", err.Error())
			log.Println(errMsg)
			logError(errMsg)
			return
		}
		fmt.Printf("Welkom: %s, Jouw kenteken is %s.\n", name, licenseplate)
	}
}

func logError(errMsg string) {
	file, err := os.OpenFile("errorlogs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Failed to open errorlogs.txt:", err.Error())
		return
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(errMsg)
}
