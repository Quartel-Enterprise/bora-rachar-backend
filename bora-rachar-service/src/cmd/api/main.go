package main

import (
	"database/sql"
	"fmt"
	"github.com/Quartel-Enterprise/Backend/bora-rachar-service/src/cmd/generated-code"
	"github.com/Quartel-Enterprise/Backend/bora-rachar-service/src/internal/infra/database"
	"log"
	"net/http"
	"os"
)

var MYSLQ *sql.DB

func init() {
	config := database.GetConfig()
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
