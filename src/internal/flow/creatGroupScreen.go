package flow

import (
	"context"
	"encoding/json"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func CreatGroupScreen(w http.ResponseWriter, r *http.Request, params swagger.CreatGroupScreenParams, db *sqlx.DB) {
	userId := params.UserId

	var requestBody swagger.GroupsScreenRequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	if err := repository_query.CreatGroupWithParticipants(context.Background(), db, requestBody, userId); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	w.WriteHeader(http.StatusCreated)
}
