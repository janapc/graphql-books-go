package database

import (
	"database/sql"
	"os"
)

func Connection() (*sql.DB, error) {
	mysqlUrl := os.Getenv("MYSQL_URL")
	database, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		return nil, err
	}

	return database, err
}
