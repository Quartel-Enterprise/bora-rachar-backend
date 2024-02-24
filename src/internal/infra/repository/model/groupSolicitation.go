package repository_model

import (
	"time"
)

type GroupSolicitation struct {
	id        string
	userId    string
	groupId   string
	createdAt time.Time
}
