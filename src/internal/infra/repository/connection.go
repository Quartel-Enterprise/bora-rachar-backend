package repository

import (
	"database/sql"
	"fmt"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	"github.com/go-sql-driver/mysql"
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

func Connect() (*sql.DB, error) {
	config := getConfig()
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("MYSQL: error when starting connection. Error: %s", err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("MYSQL: error when communicating with database. Error: %s", err.Error())
	}
	logger.Info.Println("Successful connecting to database!")
	return db, nil
}
