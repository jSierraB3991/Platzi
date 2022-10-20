package server

import (
	"context"

	"github.com/jsierra3991/platzi/proto/models"
	"github.com/jsierra3991/platzi/proto/repository"
	"github.com/jsierra3991/platzi/proto/studentpb"
)

type Server struct {
	repo repository.Repository
	studentpb.UnimplementedStudenServiceServer
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	var student = &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}
	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}
