package repository_model

import (
	"time"
)

type User struct {
	Id        string
	Name      *string
	Avatar    *string
	Email     *string
	PixKey    *string    `db:"pix_key"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
