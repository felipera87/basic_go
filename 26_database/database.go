package main

import (
	"database/sql"
	"fmt"
	"log"

	// this package is imported like this because who use the driver is the sql package
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "root:superpass@/testdb?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		// this error is related to driver or other things, not the connection itself
		log.Fatal(err)
	}
	defer db.Close()

	// this will actually test if the connection succeded
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection is open")

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal("err")
	}
	defer rows.Close()

	fmt.Println(rows)
}
