package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	gcdb, err := sql.Open("sqlite", "./chat_app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer gcdb.Close()

	err = refreshDB(gcdb)
	if err != nil {
		log.Fatal(err)
	}

}
