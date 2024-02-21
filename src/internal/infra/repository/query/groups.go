package repository_query

import (
	repository_model "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/model"
	"github.com/jmoiron/sqlx"
)

func GetGroups(userId string, db *sqlx.DB) ([]repository_model.Group, error) {
	var groups []repository_model.Group

	var query = "SELECT g.* from `group` g INNER JOIN group_participant gp  ON g.id = gp.group_id WHERE  gp.user_id  = ?;"
	err := db.Unsafe().Select(&groups, query, userId)

	if err != nil {
		return nil, err
	}

	return groups, err
}

func GetParticipants[T any](groupId T, db *sqlx.DB) ([]repository_model.JoinUserAndGroupParticipant, error) {
	var participants []repository_model.JoinUserAndGroupParticipant

	var query = "SELECT * from `group_participant` g INNER JOIN `user` u  ON g.user_id = u.id where g.group_id = ?;"
	err := db.Unsafe().Select(&participants, query, groupId)

	if err != nil {
		return nil, err
	}

	return participants, nil
}
