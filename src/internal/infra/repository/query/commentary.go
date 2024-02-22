package repository_query

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func AddCommentary(ctx context.Context, userId string, commentary string, expenseId string, db *sqlx.DB) error {
	var query = "INSERT INTO expense_commentary (user_id, expense_id, commentary, created_at, updated_at) value(?, ?, ?, ?, ?)"
	_, err := db.ExecContext(
		ctx,
		query,
		userId, expenseId, commentary, time.DateTime, time.DateTime,
	)

	if err != nil {
		return fmt.Errorf("error when inserting comment: %s", err.Error())
	}

	return nil
}

func UpdateCommentary(ctx context.Context, commentary string, commentaryId string, userId string, db *sqlx.DB) error {
	var query = "UPDATE expense_commentary SET user_id=?, commentary=?, updated_at=? WHERE id=?;"

	_, err := db.ExecContext(
		ctx,
		query,
		userId, commentary, time.DateTime, commentaryId,
	)

	if err != nil {
		return fmt.Errorf("error when updating comment: %s", err.Error())
	}

	return nil
}

func DeleteCommentary(ctx context.Context, id string, db *sqlx.DB) error {
	var query = "UPDATE expense_commentary SET updated_at=?, deleted_at=? WHERE id=?;"

	_, err := db.ExecContext(
		ctx,
		query,
		time.DateTime, id,
	)

	if err != nil {
		return fmt.Errorf("error when deleting comment: %s", err.Error())
	}

	return nil
}
