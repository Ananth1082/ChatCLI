package server

import (
	"log"
	"net"
)

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
