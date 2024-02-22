package flow

import (
	"context"
	"encoding/json"
	swagger "github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	repository_model "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/model"
	"github.com/jmoiron/sqlx"

	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"net/http"
	"time"
)

func PostScreensExpenses(w http.ResponseWriter, r *http.Request, params swagger.PostScreensExpensesParams, db *sqlx.DB) {
	var userId = params.UserId
	var requestBody swagger.ExpenseScreenRequestBody
	var expense repository_model.Expense
	var paymentSplit = make([]repository_model.ExpensePaymentSplit, len(requestBody.Participants))

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	expense = mapToExpense(requestBody)

	for _, i := range requestBody.Participants {
		if *i.UserId != requestBody.Payer {
			var value float32
			if requestBody.DivisionType == swagger.EQUALY {
				value = requestBody.Value / float32(len(requestBody.Participants))
			} else {
				value = *i.Value
			}
			var split = mapToExpenseSplit(*i.UserId, value)
			paymentSplit = append(paymentSplit, split)
		}
	}

	if err := repository_query.CreatExpenseWithPaymentSplit(context.Background(), db, userId, expense, paymentSplit); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	util.HttpResponse(w, http.StatusCreated, nil)
}

func mapToExpenseSplit(userId string, value float32) repository_model.ExpensePaymentSplit {
	return repository_model.ExpensePaymentSplit{
		UserId: userId,
		Value:  value,
	}
}

func mapToExpense(expenseScreenRequestBody swagger.ExpenseScreenRequestBody) repository_model.Expense {
	var expenseDate time.Time
	if expenseScreenRequestBody.Date != nil {
		expenseDate = time.Now()
	} else {
		expenseDate = expenseScreenRequestBody.Date.Time
	}
	return repository_model.Expense{
		GroupId:     util.NilIfStringEmpty(expenseScreenRequestBody.GroupId),
		Title:       expenseScreenRequestBody.Name,
		Description: expenseScreenRequestBody.Description,
		Category:    expenseScreenRequestBody.Category.String(),
		Avatar:      util.NilIfStringEmpty(expenseScreenRequestBody.Avatar),
		Payer:       expenseScreenRequestBody.Payer,
		Value:       expenseScreenRequestBody.Value,
		ExpenseDate: expenseDate,
		CreatedBy:   expenseScreenRequestBody.Payer,
	}
}
