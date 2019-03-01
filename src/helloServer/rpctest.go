package main

import (
	"context"
	"log"
	"net"
	"helloworld"
	"google.golang.org/grpc"
)

type Server struct{

}

func (s *Server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error){
	return &helloworld.HelloReply{Message:"hello" + in.Name},nil
}

func main(){
	l,err := net.Listen("tcp",":50051")
	if err != nil {
		log.Fatal("something wrong")
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s,&Server{})
	s.Serve(l)
}