package repository_model

import "time"

type JoinExpensePaymentSplitAndExpense struct {
	ExpenseId             string    `db:"expense_id"`
	GroupId               *string   `db:"group_id"`
	Title                 string    `db:"title"`
	Description           *string   `db:"description"`
	Category              string    `db:"category"`
	Avatar                *string   `db:"avatar"`
	ExpenseValue          float32   `db:"expense_value"`
	Payer                 string    `db:"payer"`
	ExpenseDate           time.Time `db:"expense_date"`
	ExpenseCreatedAt      time.Time `db:"expense_created_at"`
	ExpensePaymentSplitId string    `db:"split_id"`
	UserId                string    `db:"user_id"`
	SplitValue            float32   `db:"split_value"`
	CreatedAt             string    `db:"created_at"`
	IsDebtSettled         bool      `db:"is_debt_settled"`
}
