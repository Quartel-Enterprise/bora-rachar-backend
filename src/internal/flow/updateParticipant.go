package flow

import (
	"context"
	"encoding/json"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type requestBody struct {
	IsAdmin *bool `json:"is_admin"`
}

func UpdateParticipant(w http.ResponseWriter, r *http.Request, groupId string, userId string, db *sqlx.DB) {
	var body requestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	var isAdmin = util.DefaultIfNull(body.IsAdmin, false)

	if err := repository_query.UpdateParticipant(context.Background(), userId, groupId, isAdmin, db); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.HttpResponse(w, http.StatusOK, nil)
}
