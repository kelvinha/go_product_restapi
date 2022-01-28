package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = ""
	database string = "golang_test"
)

var (
	connectionString = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
