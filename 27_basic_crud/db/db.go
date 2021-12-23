package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver to connect to mysql
)

// Connect opens the connection to the database
func Connect() (*sql.DB, error) {
	stringConnection := "root:superpass@/testdb?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
