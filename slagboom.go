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
    dbName := "slagboom_db"
    dbAddress := "127.0.0.1"

    // Create data source name (DSN)
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbAddress, dbName)

    // Connect to database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Ping database to ensure connection is valid
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    // Connection successful
    fmt.Println("Connected to database!")
}
