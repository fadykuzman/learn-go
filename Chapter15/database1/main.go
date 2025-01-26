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

	DBCreate := `
	CREATE TABLE public.test (
    	id integer,
    	name character varying COLLATE pg_catalog."default"
  	)
  	WITH (
    	OIDS = FALSE
  	)
	`
	_, err = db.Exec(DBCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table successfuly created")
	}
}
