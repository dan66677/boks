package main 

import (
    "fmt"
    "database/sql"
    "log"
    _ "github.com/lib/pq" 
)

const (
    host    = "localhost"
	port     = 
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

var db *sql.DB

func InitDB() {
    psq := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

    var err error
    db, err := sql.Open("postgres", psq)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database")
}