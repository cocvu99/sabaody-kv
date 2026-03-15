//go:build linux

package server

import (
	"fmt"
	"log"
	"syscall"

	iomultiplexing "github.com/cocvu99/sabaody-kv/internal/core/io_multiplexing"
)

/*
CreateListenerFD function:
- Creating a TCP Socket
- Return a File Descriptor (a number - int)
*/
func CreateListenerFD(port int) (int, error) {
	/*
		1. Creating Socket File Descriptor (FD)
		- 1st parameter: (domain) constant number represent IPv4
		- 2nd parameter: (typ) constant number represent TCP's stream of bytes continously
	*/
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return -1,
			fmt.Errorf("Error when creating socket: %v", err)
	}

	// Make the socket non-blocking immediately
	if err := syscall.SetNonblock(fd, true); err != nil {
		syscall.Close(fd)
		return -1, fmt.Errorf("Error when setting non-blocking: %v", err)
	}

	/*
		2. Define Address and Port for binding
		- The SockaddrInet4 reprents IPv4 addresses
		- address blank = IP 0.0.0.0 (listens for all interfaces)
	*/
	addr := &syscall.SockaddrInet4{
		Port: port,
	}

	// 3. Bind the newly created socket to the address.
	// syscall function is used to assign the above `fd` to this `addr`
	// If binding fails -> Close(fd) to avoid resource leakage
	if err := syscall.Bind(fd, addr); err != nil {
		syscall.Close(fd)
		return -1, fmt.Errorf("Error when binding: %v", err)
	}

	// 4. Enable Listen mode
	// backlog: (20000 ex) The max number of connections in the Accept queue
	if err := syscall.Listen(fd, 20000); err != nil {
		syscall.Close(fd)
		return -1, fmt.Errorf("Error when listening: %v", err)
	}

	log.Printf("Listener FD %d is ready to listen on port %d", fd, port)
	return int(fd), nil
}

func Start(port int) {
	// 1. Create ListenerFD
	listenerFD, err := CreateListenerFD(port)
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}

	// 2. Create a epoll - gate-keeper
	epoll, err := iomultiplexing.NewEpoll()
	if err != nil {
		log.Fatalf("Failed to create epoll: %v", err)
	}

	// 3. Task: Epoll monitor ListenerFD (gate)
	if err := epoll.Monitor(listenerFD); err != nil {
		log.Fatalf("Failed to monitor listener: %v", err)
	}

	log.Println("Event Loop is starting...")

	// 4. Main Event Loop
	for {
		events, err := epoll.Wait()
		if err != nil {
			log.Printf("Epoll wait error: %v", err)
			continue
		}

		for _, events := range events {
			if events.Fd == listenerFD {
				// Case 1. New client (visitor)
				// The Accept() function return nfd (new FD client) and the client's address
				nfd, _, err := syscall.Accept(listenerFD)
				if err != nil {
					log.Printf("Accept() function error: %v", err)
					continue
				}

				log.Printf("New client connected with FD: %v", nfd)

				// TODO: Process and Improve logic
				// Make the socket non-blocking immediately
				if err := syscall.SetNonblock(nfd, true); err != nil {
					syscall.Close(nfd)
					log.Printf("Error when setting non-blocking with client: %v", err)
					continue
				}

				epoll.Monitor(nfd)

			} else {
				log.Printf("Data received on FD: %d", events.Fd)
			}
		}
	}
}
