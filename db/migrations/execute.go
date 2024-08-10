package migrations

import (
	"database/sql"
	"os"
)

func RunSQLFile(db *sql.DB, sqlFile string) (sql.Result, error) {
	command, err := os.ReadFile(sqlFile)
	if err != nil {
		return nil, err
	}
	res, err := db.Exec(string(command))
	if err != nil {
		return nil, err
	}
	return res, nil
}
