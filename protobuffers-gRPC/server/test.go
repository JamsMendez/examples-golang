package server

import (
	"context"
	"io"
	"log"

	"github.com/jamsmendez/protobuffers-gRPC/models"
	"github.com/jamsmendez/protobuffers-gRPC/repository"
	"github.com/jamsmendez/protobuffers-gRPC/studentpb"
	"github.com/jamsmendez/protobuffers-gRPC/testpb"
)

type TestServer struct {
	repository repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repository repository.Repository) *TestServer {
	return &TestServer{repository: repository}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repository.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	testPb := testpb.Test{
		Id:   test.ID,
		Name: test.Name,
	}

	return &testPb, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := models.Test{
		ID:   req.GetId(),
		Name: req.GetName(),
	}

	err := s.repository.SetTest(ctx, &test)
	if err != nil {
		return nil, err
	}

	setTestResponse := testpb.SetTestResponse{Id: test.ID}

	return &setTestResponse, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			setQuestionResponse := testpb.SetQuestionResponse{Ok: true}
			return stream.SendAndClose(&setQuestionResponse)
		}

		if err != nil {
			return err
		}

		question := models.Question{
			ID:       msg.GetId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
			TestID:   msg.GetTestId(),
		}

		err = s.repository.SetQuestion(context.Background(), &question)
		if err != nil {
			setQuestionResponse := testpb.SetQuestionResponse{Ok: false}
			return stream.SendAndClose(&setQuestionResponse)
		}
	}
}

func (s *TestServer) EnrollementStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			setQuestionResponse := testpb.SetQuestionResponse{Ok: true}
			return stream.SendAndClose(&setQuestionResponse)

		}

		if err != nil {
			return err
		}

		enrollment := models.Enrollment{
			StudentID: msg.GetStudentId(),
			TestID:    msg.GetTestId(),
		}

		err = s.repository.SetEnrollment(context.Background(), &enrollment)
		if err != nil {
			setQuestionResponse := testpb.SetQuestionResponse{Ok: false}
			return stream.SendAndClose(&setQuestionResponse)
		}
	}
}

func (s *TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	testID := req.GetTestId()
	students, err := s.repository.GetStudentsPerTest(context.Background(), testID)
	if err != nil {
		return err
	}

	for _, student := range students {
		nStudent := studentpb.Student{
			Id: student.ID,
			Name: student.Name,
			Age: student.Age,
		}

		err := stream.Send(&nStudent)
		// time.Sleep(2 * time.Second) // To appreciate the data stream
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {
	// msg, err := stream.Recv()
	// if err != nil {
	// 	return err
	// }
	//
	// testID := msg.GetTestId()
	testID := "test1"
	questions, err := s.repository.GetQuestionsPerTest(context.Background(), testID)
	if err != nil {
		return err
	}

	i := 0
	currentQuestion := &models.Question{}
	for {
		if i < len(questions) {
			currentQuestion = questions[i]
		}

		if i <= len(questions) {
			questionToSend := testpb.Question{
				Id: currentQuestion.ID,
				Question: currentQuestion.Question,
			} 

			err := stream.Send(&questionToSend)
			if err != nil {
				return nil
			}

			i++
		}

		answer, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Println("Answer: ", answer.GetAnswer())
	}
}
