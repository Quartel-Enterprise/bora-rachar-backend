package main

import (
	"fmt"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	repository_config "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/config"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

var MYSLQ *sqlx.DB

func init() {
	if db, err := repository_config.ConnectSqlx(); err != nil {
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
