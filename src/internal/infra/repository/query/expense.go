package repository_query

import (
	"context"
	"errors"
	"fmt"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	repository_model "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/model"
	"github.com/jmoiron/sqlx"
	"time"
)

func GetTotalAmountToPay(ctx context.Context, db *sqlx.DB, userId string) (float32, error) {
	var amount float32
	var query = "SELECT COALESCE(SUM(value),0) from expense_payment_split WHERE is_debt_settled = false AND user_id = ? and deleted_at is null;"

	if err := db.GetContext(ctx, &amount, query, userId); err != nil {
		return 0, errors.New(fmt.Sprintf("error when obtaining total payable: %s", err))
	}

	return amount, nil
}

func GetTotalAmountToReceive(ctx context.Context, db *sqlx.DB, userId string) (float32, error) {
	var amount float32
	var query = "SELECT COALESCE(SUM(eps.value),0) from expense e INNER JOIN expense_payment_split eps ON e.id  = eps.expense_id where eps.is_debt_settled = FALSE AND e.payer = ? and eps.deleted_at is null;"

	if err := db.GetContext(ctx, &amount, query, userId); err != nil {
		return 0, errors.New(fmt.Sprintf("error getting total receivable: %s", err))
	}

	return amount, nil
}

func GetExpensePaymentSplitNoSettle(ctx context.Context, db *sqlx.DB, groupId string) ([]repository_model.JoinExpensePaymentSplitAndExpense, error) {
	var response []repository_model.JoinExpensePaymentSplitAndExpense
	var query = "SELECT e.id as 'expense_id', e.value as 'expense_value',e.value as 'expense_value', eps.value as 'split_value', e.*, eps.* from expense e\n" +
		"INNER JOIN expense_payment_split eps ON e.id = eps.expense_id\n" +
		"where e.group_id = ? and eps.deleted_at is null and eps.is_debt_settled = false\n" +
		"ORDER by e.created_at;"

	if err := db.Unsafe().SelectContext(ctx, &response, query, groupId); err != nil {
		return nil, errors.New(fmt.Sprintf("error getting total receivable: %s", err))
	}

	return response, nil
}

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

func UpdateExpenseWithPaymentSplit(ctx context.Context, expense repository_model.Expense, paymentSplit []repository_model.ExpensePaymentSplit, updateSplitPayment bool, db *sqlx.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Error.Println(err.Error())
		return fmt.Errorf("error when creating transaction: %s", err)
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var queryUpdateExpense = "UPDATE BORA_RACHAR.expense\n" +
		"SET group_id=?, title=?, description=?, category=?, avatar=?, value=?, expense_date=?, payer=?, updated_at=?\n" +
		"WHERE id=?;"

	_, err = tx.ExecContext(
		ctx,
		queryUpdateExpense,
		expense.GroupId, expense.Title, expense.Description, expense.Category, expense.Avatar, expense.Value, expense.ExpenseDate, expense.Payer, time.DateTime, expense.Id)

	if err != nil {
		logger.Error.Println(err.Error())
		return fmt.Errorf("error when updating expense: %s", err.Error())
	}

	if updateSplitPayment {
		var queryDeleteExpenseSplit = "DELETE FROM BORA_RACHAR.expense_payment_split WHERE expense_id=?;"

		_, err = tx.ExecContext(
			ctx,
			queryDeleteExpenseSplit,
			expense.Id)

		if err != nil {
			logger.Error.Println(err.Error())
			return fmt.Errorf("error when deleting expense split: %s", err)
		}

		for _, i := range paymentSplit {
			_, err := tx.ExecContext(
				ctx,
				"INSERT INTO BORA_RACHAR.`expense_payment_split`(user_id, expense_id, value, created_at, is_debt_settled) VALUES (?, ?, ?, ?, ?)",
				i.UserId, expense.Id, i.Value, time.DateTime, false)
			if err != nil {
				logger.Error.Println(err.Error())
				return fmt.Errorf("error when adding split expense: %s", err.Error())
			}
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		logger.Error.Println(err.Error())
		return fmt.Errorf("error when commiting transaction: %s", err)
	}

	return nil
}

func DeleteExpenseScreen(ctx context.Context, expenseId string, db *sqlx.DB) error {
	//soft delete expense_commentary, expense_payment_split and expense
	var querySoftDeleteExpense = "UPDATE expense e\n " +
		"LEFT JOIN expense_commentary ec on e.id = ec.expense_id\n" +
		"LEFT JOIN expense_payment_split eps on e.id = eps.expense_id\n" +
		"set e.deleted_at = ?, ec.deleted_at = ?, eps.deleted_at  = ?\n" +
		"WHERE e.id = ?;"

	_, err := db.ExecContext(
		ctx,
		querySoftDeleteExpense,
		time.DateTime, time.DateTime, time.DateTime, expenseId)

	if err != nil {
		return errors.New(fmt.Sprintf("error when making soft delete expenses: %s", err))
	}

	return nil
}

func GetExpense(ctx context.Context, expenseId string, db *sqlx.DB) (repository_model.Expense, error) {
	var expense repository_model.Expense

	var query = "SELECT * FROM expense WHERE id=?;"

	err := db.Unsafe().GetContext(ctx, &expense, query, expenseId)

	if err != nil {
		return repository_model.Expense{}, errors.New(fmt.Sprintf("error when getting expense: %s", err))
	}

	return expense, nil
}

func GetExpensePaymentSplit(ctx context.Context, expenseId string, db *sqlx.DB) ([]repository_model.ExpensePaymentSplit, error) {
	var expensePaymentSplit []repository_model.ExpensePaymentSplit

	var query = "SELECT * FROM expense_payment_split WHERE expense_id=?;"

	err := db.Unsafe().SelectContext(ctx, &expensePaymentSplit, query, expenseId)

	if err != nil {
		return []repository_model.ExpensePaymentSplit{}, errors.New(fmt.Sprintf("error when getting expense split: %s", err))
	}

	return expensePaymentSplit, nil
}
