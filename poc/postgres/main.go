package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "postgres-service"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)

	err = db.Ping()
	checkError(err)

	defer db.Close()
	fmt.Println("Successfully connected!")
}
