package models

import (
	"net"

	"github.com/google/uuid"
)

const LOBBY_SESSION = "000000000000000000000000000000"

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
	return &Session{SessionID: sessionID, ChatroomID: LOBBY_SESSION, Conn: conn, Username: username}
}
