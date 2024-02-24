package util

import (
	"encoding/json"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	"net/http"
)

func DefaultIfNull[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}

	return *value
}

func NilIfStringEmpty(value *string) *string {
	if *value == "" {
		return nil
	}

	return value
}

func HttpResponse(w http.ResponseWriter, status int, err error) {
	if err != nil {
		logger.Error.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func HttpResponseWihBody(w http.ResponseWriter, status int, body interface{}, err error) {
	if err != nil {
		logger.Error.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if body != nil {
		stringResponse, err := json.Marshal(body)
		if err != nil {
			logger.Error.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(status)
		w.Write(stringResponse)
	}
}
