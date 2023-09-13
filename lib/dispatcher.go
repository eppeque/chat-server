package lib

import (
	"sync"
)

type Dispatcher struct {
	clients map[string][]*Client
	mu      sync.Mutex
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		clients: make(map[string][]*Client),
	}
}

func (d *Dispatcher) RegisterClient(room string, client *Client) {
	d.mu.Lock()
	d.clients[room] = append(d.clients[room], client)
	d.mu.Unlock()
}

func (d *Dispatcher) DispatchMessage(room, username, message string) {
	d.mu.Lock()
	for _, client := range d.clients[room] {
		client.SendMessage(message)
	}
}
