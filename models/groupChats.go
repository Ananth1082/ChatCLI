package models

type Groupchat struct {
	Members []Session
	kill    chan struct{}
}

func NewGroupChat(member Session) *Groupchat {
	return &Groupchat{
		Members: []Session{member},
		kill:    make(chan struct{}),
	}
}

func (gc *Groupchat) KillGroupchat() {
	gc.kill <- struct{}{}
}
