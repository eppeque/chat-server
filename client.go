package main

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/eppeque/chat-server/messages"
	"github.com/eppeque/chat-server/models"
)

type Client struct {
	id     int
	conn   net.Conn
	writer *bufio.Writer
	state  models.ClientState
	auth   *models.AuthManager
}

func NewClient(id int, conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)
	auth := models.NewAuthManager()

	return &Client{id, conn, writer, models.New, auth}
}

func (c *Client) Listen() {
	scanner := bufio.NewScanner(c.conn)

	hello := messages.Hello(c.auth.Challenge())
	c.SendMessage(hello)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		log.Printf("%v - %s\n", c.conn.RemoteAddr(), line)
		c.handleLine(line)
	}
}

func (c *Client) handleLine(line string) {
	mType, err := messages.DetectType(line)

	if err != nil {
		res := messages.Err(err.Error())
		c.SendMessage(res)
	}

	if c.state == models.New {
		switch mType {
		case messages.Register:
			c.handleRegister(line)
		case messages.Login:
			c.handleLogin(line)
		}
	} else if c.state == models.AuthProcess && mType == messages.Confirm {
		c.handleConfirm(line)
	} else {
		res := messages.Ok("Not implemented yet")
		c.SendMessage(res)
	}
}

func (c *Client) handleRegister(line string) {
	username, salt, hash, err := messages.ScanRegister(line)

	if err != nil {
		res := messages.Err(err.Error())
		c.SendMessage(res)
	}

	err = c.auth.Register(username, salt, hash)

	if err != nil {
		res := messages.Err(err.Error())
		c.SendMessage(res)
	}

	c.state = models.Authenticated
}

func (c *Client) handleLogin(line string) {
	// TODO: Implement
}

func (c *Client) handleConfirm(line string) {
	// TODO: Implement
}

func (c *Client) SendMessage(message string) {
	c.writer.WriteString(message)
	c.writer.Flush()
}
