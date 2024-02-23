package repository_query

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func AddParticipant(ctx context.Context, userId string, isAdmin bool, groupId string, db *sqlx.DB) error {
	var query = "INSERT INTO BORA_RACHAR.group_participant (user_id, group_id, is_admin, admission_date) VALUES(?, ?, ?, ?);"

	_, err := db.ExecContext(
		ctx,
		query,
		userId, groupId, isAdmin, time.DateTime,
	)

	if err != nil {
		return errors.New(fmt.Sprintf("error when adding participant: %s", err))
	}

	return nil
}

func UpdateParticipant(ctx context.Context, userId string, groupId string, isAdmin bool, db *sqlx.DB) error {
	var query = "UPDATE group_participant SET is_admin=?, updated_at=? WHERE group_id=? and user_id=?;"

	_, err := db.ExecContext(
		ctx,
		query,
		isAdmin, time.DateTime, groupId, userId,
	)

	if err != nil {
		return errors.New(fmt.Sprintf("error when updating participant: %s", err))
	}

	return nil
}

func RemoveParticipant(ctx context.Context, userId string, groupId string, db *sqlx.DB) error {
	var query = "UPDATE group_participant SET deleted_at=? WHERE group_id=? and user_id=?;"

	_, err := db.ExecContext(
		ctx,
		query,
		time.DateTime, groupId, userId,
	)

	if err != nil {
		return errors.New(fmt.Sprintf("error when deleting participant: %s", err))
	}

	return nil
}
