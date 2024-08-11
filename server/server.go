package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Ananth1082/Terminal_Chat_App/db"
)

const LOBBY_ID = "000000000-0000-0000-0000-000000000000"

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
			fmt.Printf("New Client has joined!!!\n\tName: %s\n\tAddress: %s\n", client.Username, client.Conn.RemoteAddr().String())
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
			if err != nil {
				fmt.Println("error occured while logging message", err)
			}
		}
	}()

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
