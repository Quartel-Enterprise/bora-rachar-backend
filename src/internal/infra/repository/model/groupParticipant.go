package repository_model

import (
	"time"
)

type GroupParticipant struct {
	Id            string
	UserId        string     `db:"user_id"`
	GroupId       string     `db:"group_id"`
	IsAdmin       bool       `db:"is_admin"`
	AdmissionDate time.Time  `db:"admission_date"`
	RemovedAt     *time.Time `db:"removed_at"`
}
