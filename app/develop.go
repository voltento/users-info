package main

import (
	"github.com/alexcesaro/log/stdlog"
	db "github.com/voltento/users-info/app/connectors/psql_connector"
)

func main() {

	logger := stdlog.GetFromFlags()

	err, db := db.NewDatabase(&db.Config{
		User:     "users-info",
		Password: "users-info",
		Database: "users-info",
	})

	if err != nil {
		logger.Error(err.Error())
	}

	err, users := db.Users()
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Infof("%v\n", users)
}
