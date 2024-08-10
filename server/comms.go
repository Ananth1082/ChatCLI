package server

import (
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/Ananth1082/Terminal_Chat_App/constants"
	"github.com/Ananth1082/Terminal_Chat_App/models"
)

// Accept loop is created to CONCURRENTLY accept requests to connect from multiple users. The infinite loops runs to listen for request. Note the .Accept() functions waits till it receives a connection.
func (server *Server) AcceptLoop() { // Use pointer receiver
	for {
		conn, err := server.ln.Accept()
		if err != nil {
			log.Println("Error while accepting connection:", err)
			continue
		}
		//Handling each connection in a seperate goroutine so as to not block the accept loop.
		go func() {
			WriteData(conn, "Successfully connected to the server\n")
			cl, err := enterInfo(server, conn)
			if err != nil {
				log.Fatal(err)
			}
			server.CommandLoop(cl)
		}()
	}
}

func enterInfo(server *Server, conn net.Conn) (*models.Session, error) {
	WriteData(conn, "Enter your username\n")
	name, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	cl := models.NewSession(conn, name)
	server.clients <- *cl
	return cl, nil
}

func (server *Server) CommandLoop(cl *models.Session) {
	defer cl.Conn.Close()
	for {
		ListCmds(*cl)
		msg, err := ReadData(cl.Conn)

		switch err.(type) {
		case *net.OpError:
			return
		}
		if err != nil {
			continue
		}

		cmd, err := strconv.Atoi(msg)
		if err != nil {
			WriteData(cl.Conn, "Give a valid command")
			continue
		}
		switch cmd - 1 {
		case constants.READ_MESSAGES: //read-data

		case constants.WRITE_MESSAGES: //write data
			msg, err := ReadData(cl.Conn)
			switch err.(type) {
			case *net.OpError:
				return
			}
			if err != nil {
				continue
			}
			server.messages <- *models.NewMessage(cl, msg)

		case constants.CHANGE_GROUP: //change group

		case constants.LEAVE: //leave chat
			WriteData(cl.Conn, "Bye have a nice day..")
			cl.Conn.Close()

		default:
			WriteData(cl.Conn, "Please enter the right command")
		}

	}
}

func ListCmds(session models.Session) {
	WriteData(session.Conn, constants.CLEAR)

	// Command list text
	cmdList := "Enter command\n\n[1]Read messages\n[2]Type new message\n[3]Change chatroom\n[4]Leave"

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
func ReadData(conn net.Conn) (string, error) {
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading data:", err)
		return "", err
	}
	return string(buf[:n-1]), nil
}

func WriteData(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
