package main

import (
	"fmt"
	"net"
	"regis_class_go/grpc/student-grpc/handlers"
	"regis_class_go/grpc/student-grpc/repositories"
	"regis_class_go/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	studentRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}
	h, err := handlers.NewStudentHandler(studentRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterHUSTStudentServer(s,h)
	fmt.Println("Listening port 2222")
	s.Serve(listen)
}