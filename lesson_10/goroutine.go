package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is listening on locahost:8000")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Printf("Connection OK with %s", connection.RemoteAddr())
		go handleConnection(connection) // all connection concurrently
	}
}

func handleConnection(c net.Conn) {

	defer func() {
		log.Printf("Connection closed with %s", c.RemoteAddr())
		c.Close()
	}()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:03:01\n"))

		if err != nil {
			log.Printf("Error writing to %s: %v", c.RemoteAddr(), err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}
