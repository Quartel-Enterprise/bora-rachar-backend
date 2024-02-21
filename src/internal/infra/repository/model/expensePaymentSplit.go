package repository_model

import (
	"math/big"
)

type ExpensePaymentSplit struct {
	Id            big.Int
	UserId        string `db:"user_id"`
	ExpenseId     string `db:"expense_id"`
	Value         float32
	CreatedAt     string `db:"created_at"`
	IsDebtSettled bool   `db:"is_debt_settled"`
}
