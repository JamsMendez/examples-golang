package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jamsmendez/protobuffers-gRPC/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		log.Fatal("Could not connection gRPC: ", err)
	}

	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)
	DoUnary(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)
	if err != nil {
		log.Fatal("DoUnary.GetTest: ", err)
	}

	fmt.Println(res.String())
}

func DoStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "1",
			Question: "Q1",
			Answer:   "R1",
			TestId:   "T1",
		},
		{
			Id:       "2",
			Question: "Q2",
			Answer:   "R2",
			TestId:   "T1",
		},
		{
			Id:       "3",
			Question: "Q3",
			Answer:   "R3",
			TestId:   "T1",
		},
	}

	stream, err := c.SetQuestions(context.Background())
	if err != nil {
		log.Fatal("DoStreaming.SetQuestions: ", err)
	}

	for _, question := range questions {
		log.Println("DoStreaming.SetQuestion.Send: ", question.GetId())
		time.Sleep(2*time.Second) // To apparience the data stream
		stream.Send(question)
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("DoStreaming.SetQuestion.CloseAndRecv: ", err)
	}

	log.Println("DoStreaming.SetQuestions: ", msg.String())
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}
	stream, err := c.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetStudentsPerTest: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving response: %v\n", err)
		}
		log.Printf("Response from GetStudentsPerTest: %v\n", msg)
	}
}

func DoBidirectionalStreaming(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequest{
		Answer: "42",
	}
	numberOfQuestions := 6

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())
	if err != nil {
		log.Fatalf("Error while calling TakeTest: %v\n", err)
	}

	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving data from TakeTest: %v\n", err)
				break
			}
			log.Printf("Response from TakeTest: %v\n", res)
		}
		close(waitChannel)
	}()
	<-waitChannel

}

