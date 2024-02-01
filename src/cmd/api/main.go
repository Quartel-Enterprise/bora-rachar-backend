package main

import (
	"database/sql"
	"fmt"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository"
	"log"
	"net/http"
	"os"
)

var MYSLQ *sql.DB

func init() {
	config := repository.GetConfig()
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	MYSLQ = db
}

func main() {
	s := BoraRacharServer{}
	h := generated.Handler(s)
	addr := ":" + os.Getenv("SERVER_PORT")

	err := http.ListenAndServe(addr, h)

	if err != nil {
		log.Fatal(err)
		return
	}
}
