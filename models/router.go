package models

import (
	"math/rand"
	"sync"
)

type MessageRouter struct {
	rooms map[string]*Room
	mu    sync.Mutex
}

const idChars = "abcdefghijklmnopqrstuvwxyz"

var Router = NewRouter()

func NewRouter() *MessageRouter {
	rooms := make(map[string]*Room)

	return &MessageRouter{rooms: rooms}
}

func (r *MessageRouter) CreateRoom(name string, sender MessageSender) string {
	id := generateId()
	room := NewRoom(id, name, sender)

	r.mu.Lock()
	r.rooms[id] = room
	r.mu.Unlock()

	return id
}

func (r *MessageRouter) SendMessage(senderId int, id, message string) {
	r.mu.Lock()
	room := r.rooms[id]
	r.mu.Unlock()

	for _, client := range room.clients {
		if client.Id() != senderId {
			client.SendMessage(message)
		}
	}
}

func generateId() string {
	id := make([]byte, 11)
	id[3] = '-'
	id[7] = '-'

	for i := 0; i < 3; i++ {
		segment := make([]byte, 3)

		for j := 0; j < 3; j++ {
			index := rand.Intn(len(idChars))
			segment[j] = idChars[index]
		}

		start := 4 * i
		id[start] = segment[0]
		id[start+1] = segment[1]
		id[start+2] = segment[2]
	}

	return string(id)
}
