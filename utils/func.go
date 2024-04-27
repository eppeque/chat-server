package utils

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ReadFlags() int {
	port := flag.Int("p", 8080, "Defines the port number to listen to")
	flag.Parse()
	return *port
}

func HandleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutting down...")
		os.Exit(0)
	}()
}
