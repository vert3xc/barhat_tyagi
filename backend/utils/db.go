package utils

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDb() (*sql.DB, error) {
	pwd, success := os.LookupEnv("db_password")
	var connStr string
	if success {
		connStr = "user=postgres password=" + pwd + " dbname=promdb sslmode=disable"
	}
	return sql.Open("postgres", connStr)
}
