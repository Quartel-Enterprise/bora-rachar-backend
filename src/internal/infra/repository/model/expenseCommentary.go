package repository_model

import (
	"time"
)

type ExpenseCommentary struct {
	Id         string
	UserId     string `db:"user_id"`
	ExpenseId  string `db:"expense_id"`
	Commentary string
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}
