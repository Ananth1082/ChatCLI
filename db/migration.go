package main

import (
	"database/sql"

	"github.com/Ananth1082/Terminal_Chat_App/db/migrations"
)

func createDB(db *sql.DB) error {
	_, err := migrations.RunSQLFile(db, "./migrations/create_database.sql")
	return err
}

func deleteDB(db *sql.DB) error {
	_, err := migrations.RunSQLFile(db, "./migrations/delete_database.sql")
	return err
}

func refreshDB(db *sql.DB) error {
	err := deleteDB(db)
	if err != nil {
		return err
	}
	err = createDB(db)
	if err != nil {
		return err
	}
	return nil
}
