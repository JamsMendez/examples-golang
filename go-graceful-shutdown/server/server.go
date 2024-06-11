package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	server    *http.Server
	shutdown  chan struct{}
	requestWg sync.WaitGroup
}

const serverPort = 8080

func Run() {
	mux := http.NewServeMux()

	srv := &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		},
		shutdown: make(chan struct{}),
	}

	mux.HandleFunc("/", srv.handler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Printf("Starting server on :%d\n", serverPort)
		if err := srv.server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %v\n", err)
		}
	}()

	<-stop
	fmt.Println("\nShutting down gracefully...")

	close(srv.shutdown)

	srvCtx, srvCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer srvCtxCancel()

	if err := srv.server.Shutdown(srvCtx); err != nil {
		fmt.Printf("Shutdown(): %v\n", err)
	}

	srv.requestWg.Wait()
	fmt.Println("Server gracefully stopped")
}

func (srv *Server) handler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-srv.shutdown:
		http.Error(w, "Server is shutting down", http.StatusServiceUnavailable)
		return
	default:
		srv.requestWg.Add(1)
		defer srv.requestWg.Done()

		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "Hello, World!")
	}
}
