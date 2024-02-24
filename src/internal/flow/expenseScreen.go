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

func CreatExpenseScreen(w http.ResponseWriter, r *http.Request, params swagger.CreatExpenseScreenParams, db *sqlx.DB) {
	var userId = params.UserId
	var expense, paymentSplit, err = mapInputToDatabaseObject(r)

	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := repository_query.CreatExpenseWithPaymentSplit(context.Background(), db, userId, expense, paymentSplit); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.HttpResponse(w, http.StatusCreated, nil)
}

func UpdateExpenseScreen(w http.ResponseWriter, r *http.Request, expenseId string, db *sqlx.DB) {
	expense, paymentSplit, err := mapInputToDatabaseObject(r)

	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	expense.Id = expenseId

	var updateSplitPayment = false
	oldExpense, err := repository_query.GetExpense(context.Background(), expenseId, db)

	if oldExpense.Value != expense.Value {
		updateSplitPayment = true
	} else {
		oldExpenseSplit, err := repository_query.GetExpensePaymentSplit(context.Background(), expenseId, db)
		if err != nil {
			util.HttpResponse(w, http.StatusInternalServerError, err)
			return
		}

		updateSplitPayment = hasSplitChange(oldExpenseSplit, paymentSplit)
	}

	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := repository_query.UpdateExpenseWithPaymentSplit(context.Background(), expense, paymentSplit, updateSplitPayment, db); err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
	}

	util.HttpResponse(w, http.StatusCreated, nil)
}

func mapInputToDatabaseObject(r *http.Request) (repository_model.Expense, []repository_model.ExpensePaymentSplit, error) {
	var requestBody swagger.ExpenseScreenRequestBody
	var expense repository_model.Expense
	var paymentSplit = make([]repository_model.ExpensePaymentSplit, len(requestBody.Participants))

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return repository_model.Expense{}, nil, err
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

	return expense, paymentSplit, nil
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

func hasSplitChange(old []repository_model.ExpensePaymentSplit, new []repository_model.ExpensePaymentSplit) bool {
	for _, o := range old {
		for _, n := range new {
			if (o.UserId == n.UserId) && (o.Value != n.Value) {
				return true
			}
		}
	}
	return false
}
