package server

import (
	"log"
	"net"

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
			server.ReadLoop(cl)
		}()
	}
}

func enterInfo(server *Server, conn net.Conn) (*models.Session, error) {
	nameBuf := make([]byte, 100)
	WriteData(conn, "Enter your username\n")
	n, err := ReadData(conn, nameBuf)
	if err != nil {
		log.Fatal(err)
	}
	cl := models.NewSession(conn, string(nameBuf[:n]))
	server.clients <- *cl
	return cl, nil
}

func (server *Server) ReadLoop(cl *models.Session) {
	defer cl.Conn.Close()
	buf := make([]byte, 2048)
	for {
		msgLength, err := ReadData(cl.Conn, buf)
		switch err.(type) {
		case *net.OpError:
			return
		}
		if err != nil {
			continue
		}
		server.messages <- *models.NewMessage(*cl, buf[:msgLength])
	}
}

func ReadData(conn net.Conn, buf []byte) (int, error) {
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading data:", err)
		return 0, err
	}
	return n - 1, nil
}

func WriteData(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
