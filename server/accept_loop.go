package server

import (
	"log"
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
