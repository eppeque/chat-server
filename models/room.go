package models

type Room struct {
	id      string
	name    string
	clients []MessageSender
}

func NewRoom(id, name string, sender MessageSender) *Room {
	clients := []MessageSender{sender}

	return &Room{id, name, clients}
}
