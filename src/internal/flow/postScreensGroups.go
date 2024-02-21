package flow

import (
	"context"
	"encoding/json"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func PostScreensGroups(w http.ResponseWriter, r *http.Request, params swagger.PostScreensGroupsParams, db *sqlx.DB) {
	userId := params.UserId

	var requestBody swagger.GroupsScreenRequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := repository_query.CreatGroupWithParticipants(context.Background(), db, requestBody, userId); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
