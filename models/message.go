package models

// Message stores all the messages for logs
type Message struct {
	Client  Session
	Message []byte
}

func NewMessage(cl Session, message []byte) *Message {
	return &Message{Client: cl, Message: message}
}
