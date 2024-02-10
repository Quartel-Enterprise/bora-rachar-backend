package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"time"
)

func GroupFindAll(db *sql.DB) ([]Group, error) {
	var groups []Group

	rows, err := db.Query("SELECT * from `group`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var item Group
		err = rows.Scan(&item.id, &item.name, &item.avatar, &item.accessCode, &item.createdBy, &item.createdAt, &item.updatedAt, &item.deletedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, item)
	}

	return groups, err
}

func CreatGroupWithParticipants(ctx context.Context, db *sql.DB, input generated.GroupsScreenRequestBody, userId string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error when creating transaction: %s", err)
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	res, err := tx.ExecContext(
		ctx,
		"INSERT INTO BORA_RACHAR.`group`(name, avatar, access_code, created_by, created_at) VALUES(?, ?, ?, ?, ?)",
		input.Name, input.Photo, "", userId, time.DateTime)

	if err != nil {
		return fmt.Errorf("error when creating group: %s", err.Error())
	}

	groupId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting group id: %s", err.Error())
	}

	for _, i := range input.Participants {
		isAdmin := (i.IsAdmin != nil && *i.IsAdmin)
		_, err := tx.ExecContext(
			ctx,
			"INSERT INTO BORA_RACHAR.`group_participant`(user_id, group_id, is_admin, admission_date) VALUES (?, ?, ?, ?)",
			i.UserId, groupId, isAdmin, time.DateTime)
		if err != nil {
			return fmt.Errorf("error when adding participants: %s", err.Error())
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error when commiting transaction: %s", err)
	}

	return nil
}
