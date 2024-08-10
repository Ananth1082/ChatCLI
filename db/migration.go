package db

import (
	"database/sql"
	"fmt"

	"github.com/Ananth1082/Terminal_Chat_App/db/migrations"
)

func createDB(db *sql.DB) error {
	_, err := migrations.RunSQLFile(db, "../db/migrations/create_database.sql")
	return err
}

func deleteDB(db *sql.DB) error {
	_, err := migrations.RunSQLFile(db, "../db/./migrations/delete_database.sql")
	return err
}

func refreshDB(db *sql.DB) error {
	fmt.Println("Started database migration: refreshing db...")
	err := deleteDB(db)
	if err != nil {
		return err
	}
	fmt.Println("Database migration: deleted tables...")
	err = createDB(db)
	if err != nil {
		return err
	}
	fmt.Println("Database migration: created tables...")
	return nil
}
