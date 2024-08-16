package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"

	"github.com/Ananth1082/Terminal_Chat_App/constants"
	"github.com/Ananth1082/Terminal_Chat_App/db"
	"github.com/Ananth1082/Terminal_Chat_App/models"
	"github.com/Ananth1082/Terminal_Chat_App/util"
)

func (server *Server) CommandLoop(cl *models.Session) {

	defer cl.Conn.Close()
	for {
		ListCmds(*cl)
		msg, err := ReadData(cl.Conn)

		switch err.(type) {
		case *net.OpError:
			return
		}
		if err == io.EOF {
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
		case constants.READ_MESSAGES: // read-data
			feed := server.createFeed()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel() // Ensure cancel is called to release resources

			messages, err := db.GetMessageFromCG(cl.ChatroomID)
			if err != nil {
				WriteData(cl.Conn, "Error occurred while getting messages: "+err.Error())
				return
			}

			var wg sync.WaitGroup
			wg.Add(1)
			// Goroutine for listening to user input to exit read mode
			go func() {
				WriteData(cl.Conn, "Press <Enter> to leave read mode\n")
				ReadData(cl.Conn)
				cancel() // Cancel the context to signal the feed goroutine
				wg.Done()
			}()

			// Display previously fetched messages
			for _, dmsg := range messages {
				WriteData(cl.Conn, fmt.Sprintf("%s: %s\n", util.PrintColorBlock(cl.Color, dmsg.UserName), dmsg.Content))
			}

			// Goroutine for reading new messages from feed
			go func() {
				for {
					select {
					case message, ok := <-feed:
						if !ok {
							return // Feed channel is closed
						}
						WriteData(cl.Conn, fmt.Sprintf("%s: %s\n", util.PrintColorBlock(cl.Color, message.Client.Username), message.Message))
					case <-ctx.Done():
						fmt.Println("cancel")
						return // Context is canceled
					}
				}
			}()
			wg.Wait() // Wait for the user to press Enter to exit read mode

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
			WriteData(cl.Conn, "Enter chatroom name\n")
			cname, _ := ReadData(cl.Conn)
			err := db.JoinChatroom(cl, cname)
			if err != nil {
				WriteData(cl.Conn, "Error joining chatroom, "+err.Error()+"\n")
			} else {
				WriteData(cl.Conn, "Successfully joined chatroom\n")
			}
		case constants.CREATE_GROUP: //create group
			WriteData(cl.Conn, "Enter chatroom name\n")
			cname, _ := ReadData(cl.Conn)
			err := db.CreateChatroom(cname)
			if err != nil {
				WriteData(cl.Conn, "Error creating chatroom, "+err.Error()+"\n")
			}
			err = db.JoinChatroom(cl, cname)
			if err != nil {
				WriteData(cl.Conn, "Error creating chatroom, "+err.Error()+"\n")
			} else {
				WriteData(cl.Conn, "Successfully created chatroom\n")
			}
		case constants.LEAVE: //leave chat
			WriteData(cl.Conn, "Bye have a nice day..")
			cl.Conn.Close()

		default:
			WriteData(cl.Conn, "Please enter the right command\n")
		}
	}
}
