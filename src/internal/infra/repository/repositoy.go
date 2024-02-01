package repository

import (
	"database/sql"
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
