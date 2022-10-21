package main

import (
	"log"
	"net"

	"github.com/jsierra3991/platzi/proto/database"
	"github.com/jsierra3991/platzi/proto/server"
	"github.com/jsierra3991/platzi/proto/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgreRepository("postgres://postgres:root@localhost:5433/platzi_go?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewTestServer(repo)
	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	if err = s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
