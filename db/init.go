package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func init() {
	gcdb, err := sql.Open("sqlite", "../db/chat_app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer gcdb.Close()

	err = refreshDB(gcdb)
	if err != nil {
		log.Fatal(err)
	}
}
