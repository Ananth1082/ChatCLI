package db

import (
	"database/sql"
	"log"

	migrations "github.com/Ananth1082/Terminal_Chat_App/db/sql"
	_ "modernc.org/sqlite"
)

type DB struct {
	database *sql.DB
}

var db DB

func init() {
	var err error
	db.database, err = sql.Open("sqlite", "../db/chat_app.db")
	if err != nil {
		log.Fatal(err)
	}
	err = migrations.RefreshDB(db.database)
	if err != nil {
		log.Fatal(err)
	}
}
