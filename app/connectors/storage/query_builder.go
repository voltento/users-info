package storage

import "github.com/go-pg/pg/orm"

func buildUserWhereEqual(user *User, query *orm.Query) *orm.Query {
	if user == nil {
		return query
	}

	if user.UserId != 0 {
		query = query.Where(tableColumnNameUserId+" = ?", user.UserId)
	}

	return query
}
