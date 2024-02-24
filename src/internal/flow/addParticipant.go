package flow

import (
	"context"
	"encoding/json"
	swagger "github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func AddParticipant(w http.ResponseWriter, r *http.Request, groupId string, db *sqlx.DB) {
	var requestBody swagger.ParticipantRequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	var isAdmin = util.DefaultIfNull(requestBody.IsAdmin, false)

	if err := repository_query.AddParticipant(context.Background(), requestBody.UserId, isAdmin, groupId, db); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.HttpResponse(w, http.StatusOK, nil)
}
