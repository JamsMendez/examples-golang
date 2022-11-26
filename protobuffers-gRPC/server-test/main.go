package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jamsmendez/protobuffers-gRPC/database"
	"github.com/jamsmendez/protobuffers-gRPC/server"
	"github.com/jamsmendez/protobuffers-gRPC/testpb"
)

func main() {
	listener, err := net.Listen("tcp", ":5061")
	if err != nil {
		log.Fatal("ServerStudent.Listen: ", err)
	}

	urlConn := "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"
	repository, err := database.NewPostgresRepository(urlConn)
	if err != nil {
		log.Fatal("ServerTest.NewPostgresRepository: ", err)
	}	

	serverTest := server.NewTestServer(repository)

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, serverTest)

	reflection.Register(s)

	fmt.Println("Server test running ...")

	if err := s.Serve(listener); err != nil {
		log.Fatal("Servertudent.Serve: ", err)
	}
}
