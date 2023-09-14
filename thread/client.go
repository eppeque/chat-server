package thread

import (
	"net"
	"strings"

	"github.com/eppeque/chat-server/msg"
	"github.com/fatih/color"
)

type Client struct {
	id         int
	conn       net.Conn
	dispatcher *Dispatcher
	username   string
	room       string
}

func NewClient(id int, conn net.Conn, dispatcher *Dispatcher) *Client {
	return &Client{
		id:         id,
		conn:       conn,
		dispatcher: dispatcher,
	}
}

func (c *Client) Listen() {
	for {
		bytes := make([]byte, 500)
		n, err := c.conn.Read(bytes)

		if err != nil {
			break
		}

		message := parseMessage(bytes, n)

		if len(message) == 0 {
			continue
		}

		color.Blue("[%v] - %v", c.conn.RemoteAddr().String(), message)

		if len(c.username) == 0 {
			c.handleRegister(message)
		} else {
			c.handleMsgs(message)
		}
	}
}

func parseMessage(bytes []byte, n int) string {
	bytes = bytes[:n]
	message := string(bytes)
	return strings.TrimSpace(message)
}

func (c *Client) handleRegister(message string) {
	username, room, err := msg.Register(message)

	if err != nil {
		errMessage := msg.Err(err.Error())
		c.SendMessage(errMessage)
		return
	}

	c.username = username
	c.room = room

	okMessage := msg.Ok("You're now registered!")
	c.SendMessage(okMessage)
	c.dispatcher.RegisterClient(room, c)

}

func (c *Client) handleMsgs(message string) {
	content, err := msg.Msgs(message)

	if err != nil {
		errMessage := msg.Err(err.Error())
		c.SendMessage(errMessage)
		return
	}

	message = msg.Msg(c.username, content)
	c.dispatcher.DispatchMessage(message, c)
}

func (c *Client) SendMessage(message string) {
	c.conn.Write([]byte(message))
}
