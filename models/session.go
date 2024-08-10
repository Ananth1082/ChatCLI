package models

import (
	"net"

	"github.com/Ananth1082/Terminal_Chat_App/constants"
	"github.com/google/uuid"
)

// Client represents the client meta data and address
type Session struct {
	ChatroomID string
	SessionID  string
	Conn       net.Conn
	Username   string
	Color      string
}

func NewSession(conn net.Conn, username string) *Session {
	sessionID := uuid.New().String()
	return &Session{SessionID: sessionID, ChatroomID: constants.LOBBY_ID, Conn: conn, Username: username}
}
