package server

import (
	"context"
	"io"
	"time"

	"github.com/jsierra3991/platzi/proto/models"
	"github.com/jsierra3991/platzi/proto/repository"
	"github.com/jsierra3991/platzi/proto/studentpb"
	"github.com/jsierra3991/platzi/proto/testpb"
)

type ServerTest struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *ServerTest {
	return &ServerTest{repo: repo}
}

func (s *ServerTest) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *ServerTest) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	var test = &models.Test{
		Id:   req.Id,
		Name: req.Name,
	}

	err := s.repo.SetTest(ctx, test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{Id: test.Id}, nil
}

func (s *ServerTest) SetQuestion(stream testpb.TestService_SetQuestionServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}

		question := &models.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}
		err = s.repo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

func (s *ServerTest) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}

		enrollment := &models.Enrollment{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}
		err = s.repo.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

func (s *ServerTest) GetStudentPerTest(req *testpb.GetStudentPerTestRequest, stream testpb.TestService_GetStudentPerTestServer) error {
	students, err := s.repo.GetStudentPerTest(context.Background(), req.GetTestId())
	if err != nil {
		return err
	}

	for _, student := range students {
		studentProto := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err = stream.Send(studentProto)
		time.Sleep(2 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
