package repository_query

import (
	"context"
	"errors"
	"fmt"
	swagger "github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"github.com/jmoiron/sqlx"
	"time"
)

func CreatGroupWithParticipants(ctx context.Context, db *sqlx.DB, input swagger.GroupsScreenRequestBody, userId string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("error when creating transaction: %s", err))
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	res, err := tx.ExecContext(
		ctx,
		"INSERT INTO BORA_RACHAR.`group`(name, avatar, access_code, created_by, created_at) VALUES(?, ?, ?, ?, ?)",
		input.Name, input.Photo, "", userId, time.DateTime)

	if err != nil {
		return errors.New(fmt.Sprintf("error when creating group: %s", err))
	}

	groupId, err := res.LastInsertId()
	if err != nil {
		return errors.New(fmt.Sprintf("error getting group id: %s", err))
	}

	for _, i := range input.Participants {
		isAdmin := i.IsAdmin != nil && *i.IsAdmin
		_, err := tx.ExecContext(
			ctx,
			"INSERT INTO BORA_RACHAR.`group_participant`(user_id, group_id, is_admin, admission_date) VALUES (?, ?, ?, ?)",
			i.UserId, groupId, isAdmin, time.DateTime)
		if err != nil {
			return errors.New(fmt.Sprintf("error when adding participants: %s", err))
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return errors.New(fmt.Sprintf("error when commiting transaction: %s", err))
	}

	return nil
}
