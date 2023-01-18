package main

import (
	"context"

	pb "github.com/Priyansh-jsk/grpcdemo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello world, I'm not ChatGPT!",
	}, nil
}
