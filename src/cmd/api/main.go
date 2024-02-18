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
	if db, err := repository.Connect(); err != nil {
		log.Fatal(err)
	} else {
		MYSLQ = db
	}
}

func main() {
	server := BoraRacharServer{}
	handler := swagger.Handler(server)
	addr := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
