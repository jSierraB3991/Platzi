package server

import (
	"context"

	"github.com/jsierra3991/platzi/proto/models"
	"github.com/jsierra3991/platzi/proto/repository"
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
