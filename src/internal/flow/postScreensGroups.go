package flow

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository"
	"net/http"
)

func PostScreensGroups(w http.ResponseWriter, r *http.Request, params swagger.PostScreensGroupsParams, db *sql.DB) {
	userId := params.UserId

	var requestBody swagger.GroupsScreenRequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := repository.CreatGroupWithParticipants(context.Background(), db, requestBody, userId); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
