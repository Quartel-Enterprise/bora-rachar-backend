package main

import (
	"github.com/Quartel-Enterprise/Backend/bora-rachar-service/src/cmd/generated-code"
	"net/http"
)

func main() {
	s := BoraRacharServer{}
	h := generated.Handler(s)

	http.ListenAndServe(":8080", h)
}
