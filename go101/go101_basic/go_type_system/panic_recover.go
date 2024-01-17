package main

import (
	"errors"
	"log"
	"net"
)

// Avoid panics crashing programs

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// Handle each client connection in a new goroutine
		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("capture a panic:", v)
			log.Println("avoid crashing the program")
		}
		c.Close()
	}()
	panic(errors.New("just a demo"))
}
