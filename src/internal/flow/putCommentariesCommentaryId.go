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

func PutCommentariesCommentaryId(w http.ResponseWriter, r *http.Request, commentaryId string, db *sqlx.DB) {
	var requestBody swagger.ExpenseCommentRequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	if err := repository_query.UpdateCommentary(context.Background(), requestBody.Commentary, commentaryId, requestBody.UserId, db); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	util.HttpResponse(w, http.StatusCreated, nil)
}
