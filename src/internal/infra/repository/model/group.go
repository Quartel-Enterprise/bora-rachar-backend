package repository_model

import (
	"database/sql"
	"time"
)

type Group struct {
	Id         string
	Name       string
	Avatar     string
	AccessCode sql.NullString `db:"access_code"`
	CreatedBy  string         `db:"created_by"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  *time.Time     `db:"updated_at"`
	DeletedAt  *time.Time     `db:"deleted_at"`
}
