package server

import (
	"fmt"
	"log"
	"syscall"
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
	return fd, nil
}
