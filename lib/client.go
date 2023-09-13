package lib

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
	id   int
}

func NewClient(conn net.Conn, id int) *Client {
	return &Client{
		conn: conn,
		id:   id,
	}
}

func (c *Client) Listen(dispatcher *Dispatcher) {
	username := ""
	room := ""

	for {
		bytes := make([]byte, 1024)
		n, err := c.conn.Read(bytes)

		if err != nil {
			break
		}

		bytes = bytes[:n]
		message := string(bytes)
		message = strings.TrimSpace(message)

		if len(message) != 0 {
			fmt.Printf("[%v] - %v\n", c.conn.RemoteAddr().String(), message)
		}

		if len(username) == 0 {
			// Handle +REGISTER...
			dispatcher.RegisterClient(room, c)
		} else {
			// Handle +MSGS...
			dispatcher.DispatchMessage(room, username, message)
		}
	}
}

func (c *Client) SendMessage(message string) {
	c.conn.Write([]byte(message))
}
