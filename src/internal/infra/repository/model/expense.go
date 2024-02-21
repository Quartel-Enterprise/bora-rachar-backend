package repository_model

import (
	"time"
)

type Expense struct {
	Id          string
	GroupId     *string `db:"group_id"`
	Title       string
	Description *string
	Category    string
	Avatar      *string
	Value       float32
	Payer       string
	ExpenseDate time.Time  `db:"expense_date"`
	CreatedBy   string     `db:"created_by"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
