package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/eppeque/chat-server/lib"
)

type ThreadCounter struct {
	v  int
	mu sync.Mutex
}

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Couldn't start the server", err)
		os.Exit(1)
	}

	fmt.Println("Server is listening on port 8080...")

	dispatcher := lib.NewDispatcher()
	idCounter := 0
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Couldn't accept TCP connection", err)
			os.Exit(1)
		}

		go handleConnection(conn, idCounter, dispatcher)
		idCounter++
	}
}

func handleConnection(conn net.Conn, id int, dispatcher *lib.Dispatcher) {
	fmt.Printf("Thread #%d created\n", id)

	client := lib.NewClient(conn, id)
	client.Listen(dispatcher)

	fmt.Printf("Thread #%d terminated\n", id)
}
