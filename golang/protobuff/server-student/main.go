package main

import (
	"log"
	"net"

	"github.com/jsierra3991/platzi/proto/database"
	"github.com/jsierra3991/platzi/proto/server"
	"github.com/jsierra3991/platzi/proto/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgreRepository("postgres://postgres:root@localhost:5432/platzi_go?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewStudentServer(repo)
	s := grpc.NewServer()
	studentpb.RegisterStudenServiceServer(s, server)

	reflection.Register(s)

	if err = s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
