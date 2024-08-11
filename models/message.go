package models

// Message stores all the messages for logs
type Message struct {
	Client  *Session
	Message string
}

type DisplayMessage struct {
	UserName  string `db:"user_name"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}

func NewMessage(cl *Session, message string) *Message {
	return &Message{Client: cl, Message: message}
}
