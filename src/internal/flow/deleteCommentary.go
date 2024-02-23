package flow

import (
	"context"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func DeleteCommentary(w http.ResponseWriter, r *http.Request, commentaryId string, db *sqlx.DB) {
	if err := repository_query.DeleteCommentary(context.Background(), commentaryId, db); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	util.HttpResponse(w, http.StatusCreated, nil)
}
