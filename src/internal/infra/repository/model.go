package repository

import (
	"database/sql"
	"math/big"
	"time"
)

type User struct {
	id        string
	userId    string
	name      string
	email     string
	pixKey    sql.NullString
	createdAt time.Time
	updatedAt sql.NullTime
	deletedAt sql.NullTime
}

type GroupParticipant struct {
	id            string
	userId        string
	groupId       string
	isAdmin       bool
	admissionDate time.Time
	removedAt     sql.NullTime
}

type GroupSolicitation struct {
	id        string
	userId    string
	groupId   string
	createdAt time.Time
}

type Group struct {
	id         string
	name       string
	avatar     string
	accessCode string
	createdBy  string
	createdAt  time.Time
	updatedAt  sql.NullTime
	deletedAt  sql.NullTime
}

type ExpensePaymentSplit struct {
	id              string
	userId          string
	expenseId       string
	value           big.Float
	transactionType string
	createdAt       time.Time
	isDebtSettled   bool
}

type ExpenseCommentary struct {
	id         string
	userId     string
	expenseId  string
	commentary string
	createdAt  time.Time
	updatedAt  sql.NullTime
	deletedAt  sql.NullTime
}

type Expense struct {
	id          string
	groupId     string
	title       string
	description string
	category    string
	avatar      string
	value       big.Float
	expenseDate time.Time
	createdBy   string
	createdAt   time.Time
	updatedAt   sql.NullTime
	deletedAt   sql.NullTime
}