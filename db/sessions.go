package db

import (
	"database/sql"
	"log"

	"github.com/Ananth1082/Terminal_Chat_App/models"
)

func IsUniqueUsername(name string) bool {
	query := `
	SELECT user_name FROM sessions WHERE user_name = ?
	`
	row := db.database.QueryRow(query, name)
	var userName string
	err := row.Scan(&userName)

	if err == sql.ErrNoRows {
		// Username is unique (not found in the database)
		return true
	} else if err != nil {
		// Some other error occurred, handle it appropriately
		log.Println("Error querying database:", err)
		return false
	}

	// Username already exists
	return false
}

func LogSession(session *models.Session) error {

	query := `
	INSERT INTO sessions (user_name,local_IP ,created_at, last_active_at)
	VALUES ( ?, ?,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
`
	_, err := db.database.Exec(query, session.Username, session.Conn.RemoteAddr().String())

	if err != nil {
		return err
	}
	return nil
}

func LeaveSession(session models.Session) error {
	query := `
	UPDATE sessions SET last_active_at=CURRENT_TIMESTAMP WHERE user_name = ?
	`
	_, err := db.database.Exec(query, session.Username)
	if err != nil {
		return err
	}
	return nil
}
