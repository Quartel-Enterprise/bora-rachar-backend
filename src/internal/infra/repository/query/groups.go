package repository_query

import (
	"context"
	"errors"
	"fmt"
	swagger "github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	repository_model "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/model"
	"github.com/jmoiron/sqlx"
	"time"
)

func GetGroups(userId string, db *sqlx.DB) ([]repository_model.Group, error) {
	var groups []repository_model.Group

	var query = "SELECT g.* from `group` g \n" +
		"INNER JOIN group_participant gp ON g.id = gp.group_id\n" +
		"WHERE  gp.user_id = ? and g.deleted_at is null;"

	err := db.Unsafe().Select(&groups, query, userId)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error when getting group: %s", err))
	}

	return groups, err
}

func GetParticipants(ctx context.Context, groupId string, db *sqlx.DB) ([]repository_model.JoinUserAndGroupParticipant, error) {
	var participants []repository_model.JoinUserAndGroupParticipant

	var query = "SELECT * from `group_participant` g \n" +
		"INNER JOIN `user` u  ON g.user_id = u.id\n" +
		"where g.group_id = ? and g.deleted_at is null;"
	err := db.Unsafe().SelectContext(ctx, &participants, query, groupId)

	if err != nil {
		return nil, err
	}

	return participants, nil
}

func GetGroup(ctx context.Context, groupId string, db *sqlx.DB) (repository_model.Group, error) {
	var group repository_model.Group

	var query = "SELECT * from `group` \n" +
		"WHERE  id = ? and deleted_at is null;"

	err := db.Unsafe().GetContext(ctx, &group, query, groupId)

	if err != nil {
		return repository_model.Group{}, errors.New(fmt.Sprintf("error when getting group: %s", err))
	}

	return group, err
}

func UpdateGroup(ctx context.Context, groupId string, name string, avatar string, db *sqlx.DB) error {
	var query = "UPDATE `group` SET name=?, avatar=?, updated_at=? WHERE id=? and deleted_at is null;"

	_, err := db.ExecContext(
		ctx,
		query,
		name, avatar, time.DateTime, groupId,
	)

	if err != nil {
		return errors.New(fmt.Sprintf("error when updating group info: %s", err))
	}

	return nil
}

func DeleteGroup(ctx context.Context, groupId string, db *sqlx.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("error when creating transaction: %s", err))
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	//soft delete group, group_participant and group_solicitation
	var querySoftDeleteGroup = "UPDATE `group`  g \n" +
		"LEFT JOIN group_participant gp on g.id = gp.group_id\n" +
		"LEFT JOIN group_solicitation  gs on g.id = gs.group_id\n" +
		"set g.deleted_at = ?, gp.deleted_at = ?, gs.deleted_at  = ?\n" +
		"WHERE g.id = ?;"

	_, err = tx.ExecContext(
		ctx,
		querySoftDeleteGroup,
		time.DateTime, time.DateTime, time.DateTime, groupId)

	if err != nil {
		return errors.New(fmt.Sprintf("error when making soft delete group: %s", err))
	}

	//soft delete expense_commentary, expense_payment_split and expense
	var querySoftDeleteExpense = "UPDATE expense e\n " +
		"LEFT JOIN expense_commentary ec on e.id = ec.expense_id\n" +
		"LEFT JOIN expense_payment_split eps on e.id = eps.expense_id\n" +
		"set e.deleted_at = ?, ec.deleted_at = ?, eps.deleted_at  = ?\n" +
		"WHERE e.group_id = ?;"

	_, err = tx.ExecContext(
		ctx,
		querySoftDeleteExpense,
		time.DateTime, time.DateTime, time.DateTime, groupId)

	if err != nil {
		return errors.New(fmt.Sprintf("error when making soft delete expenses: %s", err))
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return errors.New(fmt.Sprintf("error when commiting transaction: %s", err))
	}

	return nil
}

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
