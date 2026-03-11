//go:build linux

package iomultiplexing

import (
	"fmt"
	"syscall"

	"github.com/cocvu99/sabaody-kv/internal/config"
)

// Event represent gereral event OS independent
type Event struct {
	Fd int
}

/*
Epoll struct holds 3 parameter:
1. Epoll File Descriptor (FD)
2. Epoll Event: Declare to optimize memory & GC and re-use when Event Loop running
	-> Zero-Allocation
	- event: eg. read, write
	- fd to monitor
	- pad: Extra space
3. genericEvents: Abstraction (OS independent future)
*/

type Epoll struct {
	fd            int
	epollEvents   []syscall.EpollEvent
	genericEvents []Event
}

/*
NewEpoll creates a new epoll instance.
Return a pointer to the Epoll struct
*/
func NewEpoll() (*Epoll, error) {
	// 1. Create epoll instance
	epollFD, err := syscall.EpollCreate(0)
	if err != nil {
		return nil, fmt.Errorf("Error when creating epoll: %v", err)
	}

	return &Epoll{
		fd:            epollFD,
		epollEvents:   make([]syscall.EpollEvent, config.MaxConnection),
		genericEvents: make([]Event, config.MaxConnection),
	}, nil
}

func (e *Epoll) Monitor(fd int) error {
	// 2. Define the event want to monitor
	event := &syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(fd),
	}
	// 3. Register the fd to the epoll instance

	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, event)
	if err != nil {
		return fmt.Errorf("Error adding fd to epoll: %v", err)
	}

	return nil
}
