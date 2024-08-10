package db

import (
	"github.com/Ananth1082/Terminal_Chat_App/models"
)

func LogSession(session *models.Session) error {

	query := `
	INSERT INTO sessions (session_id, user_name,local_IP ,created_at, last_active_at)
	VALUES (?, ?, ?,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
`
	_, err := db.database.Exec(query, session.SessionID, session.Username, session.Conn.RemoteAddr().String())
	if err != nil {
		return err
	}
	return nil
}

func LeaveSession(session models.Session) error {
	query := `
	UPDATE sessions SET last_active_at=CURRENT_TIMESTAMP WHERE session_id = ?
	`
	_, err := db.database.Exec(query, session.SessionID)
	if err != nil {
		return err
	}
	return nil
}
