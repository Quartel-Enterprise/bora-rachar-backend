package repository_query

import (
	"context"
	"fmt"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/model"
	"github.com/jmoiron/sqlx"
	"time"
)

func CreatExpenseWithPaymentSplit(ctx context.Context, db *sqlx.DB, userId string, expense repository_model.Expense, paymentSplit []repository_model.ExpensePaymentSplit) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Error.Println(err.Error())
		return fmt.Errorf("error when creating transaction: %s", err)
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var query = "INSERT INTO BORA_RACHAR.`expense`(group_id, title, description, category, avatar, value, expense_date, payer, created_by, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := tx.ExecContext(
		ctx,
		query,
		expense.GroupId, expense.Title, expense.Description, expense.Category, expense.Avatar, expense.Value, expense.ExpenseDate, expense.Payer, userId, time.DateTime)

	if err != nil {
		logger.Error.Println(err.Error())
		return fmt.Errorf("error when creating group: %s", err.Error())
	}

	expenseId, err := res.LastInsertId()
	if err != nil {
		logger.Error.Println(err.Error())
		return fmt.Errorf("error getting expense id: %s", err.Error())
	}

	for _, i := range paymentSplit {
		_, err := tx.ExecContext(
			ctx,
			"INSERT INTO BORA_RACHAR.`expense_payment_split`(user_id, expense_id, value, created_at, is_debt_settled) VALUES (?, ?, ?, ?, ?)",
			i.UserId, expenseId, i.Value, time.DateTime, false)
		if err != nil {
			logger.Error.Println(err.Error())
			return fmt.Errorf("error when adding split expense: %s", err.Error())
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		logger.Error.Println(err.Error())
		return fmt.Errorf("error when commiting transaction: %s", err)
	}

	return nil
}
