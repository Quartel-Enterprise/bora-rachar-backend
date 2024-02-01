package repository

import (
	"database/sql"
	"math/big"
)

type User struct {
	id        string
	userId    string
	name      sql.NullString
	email     sql.NullString
	pixKey    sql.NullString
	createdAt sql.NullTime
	updatedAt sql.NullTime
	deletedAt sql.NullTime
}

type GroupParticipant struct {
	id            string
	userId        string
	groupId       string
	isAdmin       bool
	admissionDate sql.NullTime
	removedAt     sql.NullTime
}

type GroupSolicitation struct {
	id        string
	userId    string
	groupId   string
	createdAt sql.NullTime
}

type Group struct {
	id         string
	name       sql.NullString
	avatar     sql.NullString
	accessCode sql.NullString
	createdBy  sql.NullString
	createdAt  sql.NullTime
	updatedAt  sql.NullTime
	deletedAt  sql.NullTime
}

type ExpensePaymentSplit struct {
	id              string
	userId          string
	expenseId       string
	value           big.Float
	transactionType string
	createdAt       sql.NullTime
	isDebtSettled   bool
}

type ExpenseCommentary struct {
	id         string
	userId     string
	expenseId  string
	commentary sql.NullString
	createdAt  sql.NullTime
	updatedAt  sql.NullTime
	deletedAt  sql.NullTime
}

type Expense struct {
	id          string
	groupId     string
	title       sql.NullString
	description sql.NullString
	category    sql.NullString
	avatar      sql.NullString
	value       big.Float
	expenseDate sql.NullTime
	createdBy   sql.NullString
	createdAt   sql.NullTime
	updatedAt   sql.NullTime
	deletedAt   sql.NullTime
}
