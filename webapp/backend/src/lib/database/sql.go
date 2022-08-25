package database

import (
	"database/sql"

	"github.com/gaertnerl/translate-me.git/webserver/lib/env"
	_ "github.com/go-sql-driver/mysql"
)

func DBHandle() *sql.DB {

	user := env.Env.DB_USERNAME
	password := env.Env.DB_PASSWORD
	host := env.Env.DB_HOST
	port := env.Env.DB_PORT
	name := env.Env.DB_NAME

	dataSource := user + ":" + password + "@tpc(" + host + ":" + port + ")/" + name

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic("cannot create sql connection, reason: \n" + err.Error())
	}
	return db
}
