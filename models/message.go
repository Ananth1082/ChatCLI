package models

// Message stores all the messages for logs
type Message struct {
	Client  *Session
	Message string
}

func NewMessage(cl *Session, message string) *Message {
	return &Message{Client: cl, Message: message}
}
