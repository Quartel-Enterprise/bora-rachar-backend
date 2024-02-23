package flow

import (
	"context"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func DeleteGroup(w http.ResponseWriter, r *http.Request, groupId string, db *sqlx.DB) {
	if err := repository_query.DeleteGroup(context.Background(), groupId, db); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.HttpResponse(w, http.StatusOK, nil)
}
