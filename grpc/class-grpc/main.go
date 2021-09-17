package main

import (
	"fmt"
	"net"
	"regis_class_go/grpc/class-grpc/handlers"
	"regis_class_go/grpc/class-grpc/repositories"
	"regis_class_go/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	classRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}
	h, err := handlers.NewClassHandler(classRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterHUSTClassServer(s,h)
	fmt.Println("Listening port 2223")
	s.Serve(listen)
	
}