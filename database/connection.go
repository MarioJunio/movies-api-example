package database

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	url := fmt.Sprintf("postgres://postgres:bitlyuf@movies-pg:5432/postgres?sslmode=disable")
	db, error := sql.Open("postgres", url)

	CheckError(error)

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}