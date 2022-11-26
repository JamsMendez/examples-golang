package server

import (
	"context"

	"github.com/jamsmendez/protobuffers-gRPC/models"
	"github.com/jamsmendez/protobuffers-gRPC/repository"
	"github.com/jamsmendez/protobuffers-gRPC/studentpb"
	"github.com/jamsmendez/protobuffers-gRPC/testpb"
)

type StudentServer struct {
	repository repository.Repository
	studentpb.UnimplementedStudentServiceServer
	testpb.UnimplementedTestServiceServer
}

func NewStudentServer(repository repository.Repository) *StudentServer {
	return &StudentServer{repository: repository}
}

func (s *StudentServer) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repository.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	studentPb := studentpb.Student{
		Id:   student.ID,
		Name: student.Name,
		Age:  student.Age,
	}

	return &studentPb, nil
}

func (s *StudentServer) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := models.Student{
		ID:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repository.SetStudent(ctx, &student)
	if err != nil {
		return nil, err
	}

	setStudentResponse := studentpb.SetStudentResponse{Id: student.ID}

	return &setStudentResponse, nil
}

func (s *StudentServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
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

func (s *StudentServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
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
