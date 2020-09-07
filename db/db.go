package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "me"
	password = "thang1511"
	dbname   = "project"
)

func getConnect() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ := sqlx.Connect("postgres", psqlInfo)
	return db
}

var Db = getConnect()
