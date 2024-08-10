package models

import "net"

// Client represents the client meta data and address
type Client struct {
	Conn     net.Conn
	Username string
	Color    string
}

func NewClient(conn net.Conn, username string) *Client {
	return &Client{Conn: conn, Username: username}
}
