package repository_query

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetTotalAmountToPay(ctx context.Context, db *sqlx.DB, userId string) (float32, error) {
	var amount float32
	var query = "SELECT COALESCE(SUM(value),0) from expense_payment_split WHERE is_debt_settled = false AND user_id = ?;"

	if err := db.GetContext(ctx, &amount, query, userId); err != nil {
		return 0, errors.New(fmt.Sprintf("error when obtaining total payable: %s", err))
	}

	return amount, nil
}

func GetTotalAmountToReceive(ctx context.Context, db *sqlx.DB, userId string) (float32, error) {
	var amount float32
	var query = "SELECT COALESCE(SUM(eps.value),0) from expense e INNER JOIN expense_payment_split eps ON e.id  = eps.expense_id where eps.is_debt_settled = FALSE AND e.payer = ?;"

	if err := db.GetContext(ctx, &amount, query, userId); err != nil {
		return 0, errors.New(fmt.Sprintf("error getting total receivable: %s", err))
	}

	return amount, nil
}
