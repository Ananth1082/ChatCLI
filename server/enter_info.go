package server

import (
	"log"
	"net"

	"github.com/Ananth1082/Terminal_Chat_App/db"
	"github.com/Ananth1082/Terminal_Chat_App/models"
)

func enterInfo(server *Server, conn net.Conn) (*models.Session, error) {
	WriteData(conn, "Enter your username\n")
	var name string
	var err error
	for { //name loop
		name, err = ReadData(conn)
		if err != nil {
			WriteData(conn, "Error reading username. Please try again.\n")
			log.Println("Error reading username:", err)
			continue
		}

		if db.IsUniqueUsername(name) {
			WriteData(conn, "Username accepted.\n")
			break
		}

		WriteData(conn, "Username taken. Please try a different one.\n")
	}
	WriteData(conn, "Enter your favourite color\n")
	color, _ := ReadData(conn)
	cl := models.NewSession(conn, name, color)
	server.clients <- *cl
	return cl, nil
}
