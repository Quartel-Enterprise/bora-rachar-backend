package flow

import (
	"context"
	"database/sql"
	"encoding/json"
	swagger "github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"net/http"
	"time"
)

func PostScreensExpenses(w http.ResponseWriter, r *http.Request, params swagger.PostScreensExpensesParams, db *sql.DB) {
	var userId = params.UserId
	var requestBody swagger.ExpenseScreenRequestBody
	var expense repository.Expense
	var paymentSplit = make([]repository.ExpensePaymentSplit, len(requestBody.Participants))

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
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

	if err := repository.CreatExpenseWithPaymentSplit(context.Background(), db, userId, expense, paymentSplit); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func mapToExpenseSplit(userId string, value float32) repository.ExpensePaymentSplit {
	return repository.ExpensePaymentSplit{
		UserId: userId,
		Value:  value,
	}
}

func mapToExpense(expenseScreenRequestBody swagger.ExpenseScreenRequestBody) repository.Expense {
	var expenseDate time.Time
	if expenseScreenRequestBody.Date != nil {
		expenseDate = time.Now()
	} else {
		expenseDate = expenseScreenRequestBody.Date.Time
	}
	return repository.Expense{
		GroupId:     util.NilIfStringEmpty(expenseScreenRequestBody.GroupId),
		Title:       expenseScreenRequestBody.Name,
		Description: expenseScreenRequestBody.Description,
		Category:    expenseScreenRequestBody.Category.String(),
		Avatar:      util.NilIfStringEmpty(expenseScreenRequestBody.Avatar),
		Value:       expenseScreenRequestBody.Value,
		ExpenseDate: expenseDate,
		CreatedBy:   expenseScreenRequestBody.Payer,
	}
}
