package main

import (
	"fmt"
	"net"
	"os"

	"github.com/eppeque/chat-server/lib"
	"github.com/fatih/color"
)

var (
	threadCounter *lib.ThreadCounter = lib.NewThreadCounter()
	dispatcher    *lib.Dispatcher    = lib.NewDispatcher()
)

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		color.Red("Couldn't start the server", err)
		os.Exit(1)
	}

	color.Green("Server listening on port 8080...")

	id := 0
	for {
		conn, err := ln.Accept()

		if err != nil {
			color.Red("Couldn't accept TCP connection", err)
			os.Exit(1)
		}

		go handleConnection(conn, id)
		id++
	}
}

func handleConnection(conn net.Conn, id int) {
	threadCounter.Inc()
	fmt.Printf("[INFO] - Thread created. %d thread(s) running.\n", threadCounter.Value())

	client := lib.NewClient(conn, id)
	client.Listen(dispatcher)

	threadCounter.Dec()
	fmt.Printf("[INFO] - Thread terminated. %d thread(s) running.\n", threadCounter.Value())
}
