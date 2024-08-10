package migrations

import (
	"database/sql"
	"fmt"
)

func createDB(db *sql.DB) error {
	_, err := runSQLFile(db, "../db/sql/create_database.sql")
	return err
}

func deleteDB(db *sql.DB) error {
	_, err := runSQLFile(db, "../db/./sql/delete_database.sql")
	return err
}

func RefreshDB(db *sql.DB) error {
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
