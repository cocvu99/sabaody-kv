package main

import (
	"io"
	"log"
	"net"
)

// Job: element in the queue - having the connection
type Job struct {
	conn net.Conn
}

// Worker meaning thread in the pool
type Worker struct {
	id       int
	jobQueue chan Job
}

// Queue and array of Worker
type Pool struct {
	jobQueue chan Job
	workers  []*Worker
}

/*
NewPool function: Create a new Pool
n int: Pass in the number of threads want to create in the pool
create a new queue
create a new array: number of threads
*/
func NewPool(n int) *Pool {
	return &Pool{
		jobQueue: make(chan Job),
		workers:  make([]*Worker, n),
	}
}

/*
NewWorker function: Creating a new Worker
*/
func NewWorker(id int, jobQueue chan Job) *Worker {
	return &Worker{
		id:       id,
		jobQueue: jobQueue,
	}
}

/*
Start-Worker function:
go func() -> Create new goroutine/thread
for loop: to get the connection -> call the handleconnection
*/
func (w *Worker) Start() {
	go func() {
		// chanel is thread-safe and blocking
		for job := range w.jobQueue {
			log.Printf("Worker %d is processing from %s",
				w.id,
				job.conn.RemoteAddr())

			handleConnection(job.conn)
		}
	}()
}

/*
Start-Pool function
*/
func (p *Pool) Start() {
	for i := 0; i < len(p.workers); i++ {
		worker := NewWorker(i, p.jobQueue)
		p.workers[i] = worker
		worker.Start()
	}
}

func (p *Pool) AddJob(conn net.Conn) {
	p.jobQueue <- Job{conn: conn}
}

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func handleConnection(conn net.Conn) {
	/*
		defer conn.Close()
		buf := make([]byte, 1000)

		conn.Read(buf)
		time.Sleep(5 * time.Second)

		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n Phase 2 Thread-pool\r\n"))
	*/
	log.Println("Handle conn from", conn.RemoteAddr())
	for {
		cmd, err := readCommand(conn)
		if err != nil {
			conn.Close()
			log.Println("client disconnected", conn.RemoteAddr())
			if err == io.EOF {
				break
			}
		}

		if err = respond(cmd, conn); err != nil {
			log.Println("err write:", err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP-server is started. Listening at port 3000")

	defer listener.Close()

	pool := NewPool(2)
	pool.Start()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// go handleConnection(conn)
		pool.AddJob(conn)
	}

}
