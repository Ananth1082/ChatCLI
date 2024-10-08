package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Ananth1082/Terminal_Chat_App/db"
)

// Start the server by creating a tcp listener and bind it to the server struct
func (server *Server) Start() error { // Use pointer receiver
	//ln is the listener
	ln, err := net.Listen("tcp", server.address)
	if err != nil {
		return err
	}
	defer ln.Close()
	fmt.Println("Server started on", ln.Addr())
	server.ln = ln
	go server.AcceptLoop()
	//blocking channel which doesnt allow the server to close until we pass empty struct to the quit channel.
	<-server.quitch
	return nil
}

func Run() {
	server := NewServer(":8080")

	// List clients
	go func() {
		for client := range server.clients {
			fmt.Printf("New Client has joined\n\tName: %s\n\tAddress: %s\n", client.Username, client.Conn.RemoteAddr().String())
			err := db.LogSession(&client)
			if err != nil {
				fmt.Println("error occured while logging session. error: ", err)
			}
		}
	}()

	//Message Logs
	go func() {
		for msg := range server.messages {
			err := db.LogMessages(msg)
			for _, feed := range server.userFeeds {
				select {
				case feed <- msg: // Try to send the message
					// Message sent successfully
				default:
					// The channel is blocked or closed, skip sending
				}
			}
			if err != nil {
				fmt.Println("error occured while logging message", err)
			}
		}
	}()
	go func() {
		var s string
		fmt.Println("Press <Enter> to close the server")
		fmt.Scanln(&s)
		fmt.Println("Server closed. Thank you")
		server.quitch <- struct{}{}
	}()
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
