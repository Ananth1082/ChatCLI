package db

import (
	"github.com/Ananth1082/Terminal_Chat_App/models"
	"github.com/google/uuid"
)

func LogMessages(message models.Message) error {
	messageID := uuid.New().String()
	query := `
	INSERT INTO messages (message_id,chatroom_id,session_id,content,sent_at)
	VALUES (?, ?, ?, ?,CURRENT_TIMESTAMP)
`
	_, err := db.database.Exec(query, messageID, message.Client.ChatroomID, message.Client.SessionID, string(message.Message))
	if err != nil {
		return err
	}
	return nil
}

func GetMessageFromCG(chatroomID string) (map[string]string, error) {
	res := make(map[string]string)
	query := `
	SELECT s.user_name, m.content
	FROM messages AS m
	JOIN sessions AS s ON s.session_id = m.session_id
	WHERE m.chatroom_id = ?;
	`
	rows, err := db.database.Query(query, chatroomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var message struct {
			userName string
			content  string
		}
		rows.Scan(&message)
		res[message.userName] = message.content
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
