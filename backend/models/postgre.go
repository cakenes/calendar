package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "192.168.1.10"
	port     = 5432
	user     = "postgres"
	password = "postgresql"
	dbname   = "calendar"
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Successfully connected.")
}

func Setup() {
	dir, err := os.ReadDir("db")
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		query, err := os.ReadFile("db/" + file.Name())
		if err != nil {
			panic(err)
		}
		if _, err := DB.Exec(string(query)); err != nil {
			fmt.Println("No setup required.")
			return
		}
	}

	fmt.Println("Successfully setup.")
}
