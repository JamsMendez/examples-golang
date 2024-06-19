package tcp

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run() {
	s, err := NewServerTCP(":8080")
	if err != nil {
		os.Exit(1)
		return
	}

	s.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	fmt.Println("Shutting down server ...")
	s.Stop()
	fmt.Println("Server TCP stopped")
}

type ServerTCP struct {
	wg         sync.WaitGroup
	listener   net.Listener
	shutdown   chan struct{}
	connection chan net.Conn
}

func NewServerTCP(address string) (*ServerTCP, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on address %s: %w", address, err)
	}

	server := &ServerTCP{
		listener:   listener,
		shutdown:   make(chan struct{}),
		connection: make(chan net.Conn),
	}

	return server, nil
}

func (s *ServerTCP) acceptConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				continue
			}

			s.connection <- conn
		}
	}
}

func (s *ServerTCP) handleConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		case conn := <-s.connection:
			go s.handleConnection(conn)
		}
	}
}

func (s *ServerTCP) handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintf(conn, "Welcome to TCP Server!\n")
	<-time.After(5 * time.Second)
	fmt.Fprintf(conn, "Goodbye!\n")
}

func (s *ServerTCP) Start() {
	s.wg.Add(2)

	go s.acceptConnections()
	go s.handleConnections()
}

func (s *ServerTCP) Stop() {
	close(s.shutdown)
	s.listener.Close()

	done := make(chan struct{})

	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return
	case <-time.After(time.Second):
		fmt.Println("Timed out waiting for connections to finish")
		return
	}
}
