package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jamsmendez/protobuffers-gRPC/database"
	"github.com/jamsmendez/protobuffers-gRPC/server"
	"github.com/jamsmendez/protobuffers-gRPC/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal("ServerStudent.Listen: ", err)
	}

	urlConn := "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"
	repository, err := database.NewPostgresRepository(urlConn)
	if err != nil {
		log.Fatal("ServerStudent.NewPostgresRepository: ", err)
	}

	serverStudent := server.NewStudentServer(repository)

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, serverStudent)

	reflection.Register(s)

	fmt.Println("Server student running ...")

	if err := s.Serve(listener); err != nil {
		log.Fatal("ServerStudent.Serve: ", err)
	}
}
