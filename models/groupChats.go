package models

type Groupchat struct {
	Members []Client
	kill    chan struct{}
}

func NewGroupChat(member Client) *Groupchat {
	return &Groupchat{
		Members: []Client{member},
		kill:    make(chan struct{}),
	}
}

func (gc *Groupchat) KillGroupchat() {
	gc.kill <- struct{}{}
}
