package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"user=fady password=fady host=127.0.0.1 port=5432 dbname=postgres sslmode=disable",
	)
	defer db.Close()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection to Db was successful")
	}

	connectivity := db.Ping()

	if connectivity != nil {
		panic(err)
	} else {
		fmt.Println("Good to go!")
	}

	insert, err := db.Prepare("INSERT INTO test VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(2, "second")
	if err != nil {
		panic(err)
	}
	fmt.Println("The values has been successfully stored")
}
