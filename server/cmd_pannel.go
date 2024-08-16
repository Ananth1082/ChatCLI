package server

import (
	"strings"

	"github.com/Ananth1082/Terminal_Chat_App/constants"
	"github.com/Ananth1082/Terminal_Chat_App/models"
)

func ListCmds(session models.Session) {
	// WriteData(session.Conn, constants.CLEAR)

	// Command list text
	WriteData(session.Conn, session.ChatroomID+"\n")
	cmdList := "Enter command\n\n[1]Read messages\n[2]Type new message\n[3]Change chatroom\n[4]Create chatroom\n[5]Leave"

	// Split the command list into lines
	lines := strings.Split(cmdList, "\n")

	// Find the maximum length of the lines to determine box width
	maxLineLength := 0
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	// Set the box width, including padding
	padding := 2
	boxWidth := maxLineLength + padding*2 + 30

	// Create the borders
	borderTop := constants.DOUBLE_LINE_TOP_LEFT + strings.Repeat(constants.DOUBLE_LINE_HORIZONTAL, boxWidth) + constants.DOUBLE_LINE_TOP_RIGHT
	borderBottom := constants.DOUBLE_LINE_BOTTOM_LEFT + strings.Repeat(constants.DOUBLE_LINE_HORIZONTAL, boxWidth) + constants.DOUBLE_LINE_BOTTOM_RIGHT

	// Create the boxed command list
	boxedCmdList := ""
	for _, line := range lines {
		boxedCmdList += constants.DOUBLE_LINE_VERTICAL + " " + line + strings.Repeat(" ", boxWidth-len(line)-2) + " " + constants.DOUBLE_LINE_VERTICAL + "\n"
	}

	// Add an empty line if needed to fill the box
	emptyLine := constants.DOUBLE_LINE_VERTICAL + strings.Repeat(" ", boxWidth) + constants.DOUBLE_LINE_VERTICAL + "\n"
	boxedCmdList += emptyLine

	// Send the box
	WriteData(session.Conn, borderTop+"\n")
	WriteData(session.Conn, boxedCmdList)
	WriteData(session.Conn, borderBottom+"\n")
}
