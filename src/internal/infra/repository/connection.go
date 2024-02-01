package repository

import (
	"github.com/go-sql-driver/mysql"
	"os"
)

func GetConfig() mysql.Config {
	return mysql.Config{
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Net:       "tcp",
		Addr:      "localhost:" + os.Getenv("MYSQL_PORT"),
		DBName:    os.Getenv("MYSQL_DATABASE"),
		ParseTime: true,
	}
}
