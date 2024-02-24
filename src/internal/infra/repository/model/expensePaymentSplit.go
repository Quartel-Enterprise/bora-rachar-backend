package repository_model

type ExpensePaymentSplit struct {
	Id            string
	UserId        string `db:"user_id"`
	ExpenseId     string `db:"expense_id"`
	Value         float32
	CreatedAt     string `db:"created_at"`
	IsDebtSettled bool   `db:"is_debt_settled"`
}
