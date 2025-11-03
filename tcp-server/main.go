package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	// read data from client
	log.Println(conn.RemoteAddr())
	var buf []byte = make([]byte, 1000)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// Create reader to read line-by-line
	reader := bufio.NewReader(conn)

	/*
		Create a loop:
			- Re-use and maintain the connection between client and server
			- Pass data from buffer (read) client and send to the server
			- Handle all exception error cases: client closing connection, ... TODO
			- TODO: Check all exception error case and handle them
	*/
	for {
		conn.SetReadDeadline(time.Now().Add(15 * time.Second))

		message, err := reader.ReadString('\n')

		// Handle all exception error cases
		if err != nil {
			// Normal case: Client closes connection first
			if err.Error() == "EOF" {
				log.Printf("Client %s disconnected", conn.RemoteAddr())
				return
			}
		}

		message = strings.TrimSpace(message)

		// process after 500ms
		time.Sleep(500 * time.Millisecond)

		// Create response message and its format
		response := fmt.Sprintf("Echo: %s (at %s)\n",
			message,
			time.Now().Format("15:04:05"))

		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Printf("Error writing to %s: %v", conn.RemoteAddr(), err)
			return
		}

	}
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP-server is started. Listening at port 3000")

	for {
		// conn == socket == dedicated communication channel
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// create a new goroutine to handle the connection
		go handleConnection(conn)
	}
}
