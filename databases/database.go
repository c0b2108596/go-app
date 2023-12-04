package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dbuser:8728@tcp(localhost:3306)/todolist")
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")

	return db, nil
}
