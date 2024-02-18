package repository

import (
	"database/sql"
	"math/big"
	"time"
)

type User struct {
	id        big.Int
	name      string
	email     string
	pixKey    string
	createdAt time.Time
	updatedAt *time.Time
	deletedAt *time.Time
}

type GroupParticipant struct {
	id            big.Int
	userId        string
	groupId       string
	isAdmin       bool
	admissionDate time.Time
	removedAt     *time.Time
}

type GroupSolicitation struct {
	id        big.Int
	userId    string
	groupId   string
	createdAt time.Time
}

type Group struct {
	id         big.Int
	name       string
	avatar     string
	accessCode sql.NullString
	createdBy  string
	createdAt  *time.Time
	updatedAt  *time.Time
	deletedAt  *time.Time
}

type ExpensePaymentSplit struct {
	Id            big.Int
	UserId        string
	ExpenseId     string
	Value         float32
	CreatedAt     *time.Time
	IsDebtSettled bool
}

type TransactionType string

type ExpenseCommentary struct {
	id         big.Int
	userId     string
	expenseId  string
	commentary string
	createdAt  *time.Time
	updatedAt  *time.Time
	deletedAt  *time.Time
}

type Expense struct {
	Id          big.Int
	GroupId     *string
	Title       string
	Description *string
	Category    string
	Avatar      *string
	Value       float32
	ExpenseDate time.Time
	CreatedBy   string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
