package main

import (
	"fmt"
	"log"
	"net"

	"github.com/eppeque/chat-server/utils"
)

var counter = utils.NewThreadCounter()

func main() {
	utils.HandleInterrupt()
	port := utils.ReadFlags()
	address := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln("Startup error: ", err)
	}

	log.Printf("Listening on port %d...\n", port)
	id := 0

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go handleConnection(id, conn)
		id++
	}
}

func handleConnection(id int, conn net.Conn) {
	counter.Increment()
	log.Printf("%v - NEW CLIENT - %d thread(s) running...\n", conn.RemoteAddr(), counter.Value())

	NewClient(id, conn).Listen()

	counter.Decrement()
	log.Printf("%v - DISCONNECTED - %d thread(s) running...\n", conn.RemoteAddr(), counter.Value())
}
