package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	// read buffer from client after enter is hit
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("client left..")
		conn.Close()

		// escape recursion
		return
	}

	// convert bytes from buffer to string
	message := string(bufferBytes)
	// get the remote address of the client
	clientAddr := conn.RemoteAddr().String()
	// format a response
	response := fmt.Sprintf(message + " from " + clientAddr + "\n")

	// have server print out important information
	log.Println(response)

	// let the client know what happened
	conn.Write([]byte("you sent: " + response))

	// recursive func to handle io.EOF for random disconnects
	handleConnection(conn)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to listen", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
		}
		go handleConnection(conn)
	}
}
