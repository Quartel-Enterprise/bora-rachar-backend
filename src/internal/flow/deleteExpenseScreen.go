package flow

import (
	"context"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func DeleteExpenseScreen(w http.ResponseWriter, r *http.Request, expenseId string, db *sqlx.DB) {
	if err := repository_query.DeleteExpenseScreen(context.Background(), expenseId, db); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.HttpResponse(w, http.StatusCreated, nil)
}
