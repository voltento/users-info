package database

import "github.com/go-pg/pg/orm"

func buildUserWhereEqual(user *User, query *orm.Query) *orm.Query {
	if user == nil {
		return query
	}

	if user.UserId != 0 {
		query = query.Where(tableColumnNameUserId+" = ?", user.UserId)
	}

	if len(user.FirstName) > 0 {
		query = query.Where(tableColumnNameFirstName+" = ?", user.FirstName)
	}

	if len(user.LastName) > 0 {
		query = query.Where(tableColumnNameLastName+" = ?", user.LastName)
	}

	if len(user.CountryCode) > 0 {
		query = query.Where(tableColumnNameCountryCode+" = ?", user.CountryCode)
	}

	return query
}
