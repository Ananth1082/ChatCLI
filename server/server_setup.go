package server

import (
	"net"

	"github.com/Ananth1082/Terminal_Chat_App/models"
)

type Server struct {
	address   string              //Port for localhost.
	ln        net.Listener        //Listener to listen for requests.
	quitch    chan struct{}       //quit channel to end the server. Empty struct because it takes no memory.
	clients   chan models.Session //client channel to print client as soon as they log.
	messages  chan models.Message //message channel to print message as soon as they appear.
	userFeeds []chan models.Message
}

func NewServer(addr string) *Server {
	return &Server{
		address:   addr,
		quitch:    make(chan struct{}),
		clients:   make(chan models.Session, 10), //initialize the client channel to hold 10 conncurrent sign ups.
		messages:  make(chan models.Message, 10), // Initialize the messages channel to hold 10 concurrent messages.
		userFeeds: make([]chan models.Message, 0, 20),
	}
}

func (s *Server) createFeed() chan models.Message {
	feed := make(chan models.Message)
	s.userFeeds = append(s.userFeeds, feed)
	return feed
}
