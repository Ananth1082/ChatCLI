package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Ananth1082/Terminal_Chat_App/constants"
	"github.com/Ananth1082/Terminal_Chat_App/models"
)

func CreateChatroom(chatroomName string) error {
	if !IsUniqueChatroomName(chatroomName) {
		return errors.New("chatroom name already taken")
	}
	query := `
	INSERT INTO chatrooms(chatroom_name,created_at)
	VALUES(?,CURRENT_TIMESTAMP)
	`
	_, err := db.database.Exec(query, chatroomName)
	if err != nil {
		return err
	}
	return nil
}

func IsUniqueChatroomName(chatroomName string) bool {
	query := `
		SELECT chatroom_name FROM chatrooms WHERE chatroom_name = ?
		`
	row := db.database.QueryRow(query, chatroomName)
	var chatroom string
	err := row.Scan(&chatroom)

	if err == sql.ErrNoRows {
		return true
	} else if err != nil {
		// Some other error occurred, handle it appropriately
		log.Println("Error querying database:", err)
		return false
	}

	// chatroom already exists
	return false
}

func JoinChatroom(session *models.Session, chatroomName string) error {
	query := `
	INSERT INTO members(chatroom_id,session_id,joined_at) 
	VALUES(?,?,CURRENT_TIMESTAMP)
	`
	_, err := db.database.Exec(query, chatroomName, session.Username)
	if err != nil {
		return err
	}
	session.ChangeChatroom(chatroomName)
	return nil
}

func LeaveChatroom(session *models.Session) error {
	query := `
		DELETE FROM members WHERE chatroom_id = ? AND session_id = ?
	`
	_, err := db.database.Exec(query, session.ChatroomID, session.Username)
	if err != nil {
		return err
	}
	session.ChangeChatroom(constants.LOBBY_ID)
	return nil
}
