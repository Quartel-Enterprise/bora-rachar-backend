package repository_config

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

func getConfig() mysql.Config {
	return mysql.Config{
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Net:       "tcp",
		Addr:      "localhost:" + os.Getenv("MYSQL_PORT"),
		DBName:    os.Getenv("MYSQL_DATABASE"),
		ParseTime: true,
	}
}

func ConnectSqlx() (*sqlx.DB, error) {
	config := getConfig()
	db, err := sqlx.Connect("mysql", config.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("MYSQL: error when starting connection. Error: %s", err.Error())
	}

	return db, nil
}
