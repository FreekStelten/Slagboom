package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set up database connection parameters
	dbUser := "Admin"
	dbPass := "Fonteyn@DB"
	dbName := "slagboomm_db"
	dbAddress := "127.0.0.1"

	// Create data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbAddress, dbName)

	// Connect to database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("er kan geen connecting naar de database gemaakt worden:", err.Error())
		return
	}
	defer db.Close()

	// Ping database to ensure connection is valid
	err = db.Ping()
	if err != nil {
		fmt.Println("Er kan niet gepingt worden naar de database:", err.Error())
		return
	}

	// Connection successful
	fmt.Println("Connected to database!")
}
