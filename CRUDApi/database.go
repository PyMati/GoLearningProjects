package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const defaultId = "id serial"

type Table struct {
	Name string `json:"name"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "poradnia2"
	dbname   = "golangproject"
)

func connectToDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}

func closeConnection(db *sql.DB) {
	defer db.Close()
}
