package main

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/eppeque/chat-server/models"
)

type Client struct {
	id      int
	conn    net.Conn
	writer  *bufio.Writer
	handler *models.MessageHandler
}

func NewClient(id int, conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)

	return &Client{id, conn, writer, nil}
}

func (c *Client) Listen() {
	c.handler = models.NewHandler(c)
	scanner := bufio.NewScanner(c.conn)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		log.Printf("%v - %s\n", c.conn.RemoteAddr(), line)
		c.handler.HandleLine(line)
	}
}

func (c *Client) SendMessage(message string) {
	c.writer.WriteString(message)
	c.writer.Flush()
}
