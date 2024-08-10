package models

// Message stores all the messages for logs
type Message struct {
	Client  Client
	Message []byte
}

func NewMessage(cl Client, message []byte) *Message {
	return &Message{Client: cl, Message: message}
}
