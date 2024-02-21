package repository_model

import (
	"math/big"
	"time"
)

type GroupSolicitation struct {
	id        big.Int
	userId    string
	groupId   string
	createdAt time.Time
}
