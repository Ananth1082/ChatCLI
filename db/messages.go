package db

import (
	"github.com/Ananth1082/Terminal_Chat_App/models"
	"github.com/google/uuid"
)

func LogMessages(message models.Message) error {
	messageID := uuid.New().String()
	query := `
	INSERT INTO messages (message_id,chatroom_name,user_name,content,sent_at)
	VALUES (?, ?, ?, ?,CURRENT_TIMESTAMP)
`
	db.Lock()
	_, err := db.database.Exec(query, messageID, message.Client.ChatroomID, message.Client.Username, string(message.Message))
	db.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func GetMessageFromCG(chatroom string) ([]models.DisplayMessage, error) {
	res := make([]models.DisplayMessage, 0, 50)
	query := `
	SELECT user_name, content, sent_at
	FROM messages 
	WHERE chatroom_name = ?;
	`
	rows, err := db.database.Query(query, chatroom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var message models.DisplayMessage
		err := rows.Scan(&message.UserName, &message.Content, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
