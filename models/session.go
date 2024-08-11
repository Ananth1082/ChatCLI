package models

import (
	"net"
	"strings"

	"github.com/Ananth1082/Terminal_Chat_App/constants"
	"github.com/Ananth1082/Terminal_Chat_App/util"
)

// Client represents the client meta data and address
type Session struct {
	ChatroomID string
	Conn       net.Conn
	Username   string
	Color      string
}

func NewSession(conn net.Conn, username string, favColor string) *Session {
	var color string
	favColor = strings.ToLower(favColor)
	for color = range util.ColorMap {
		if favColor == strings.ToLower(color) {
			break
		}
	}

	return &Session{ChatroomID: constants.LOBBY_ID, Conn: conn, Username: username, Color: color}
}

func (s *Session) ChangeChatroom(newChatroom string) {
	s.ChatroomID = newChatroom
}
