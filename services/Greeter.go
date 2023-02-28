package services

import (
	pb "ad-service-rpc/pb"
	"context"
	"log"
)

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func (g *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received params: %v", in.GetName())
	return &pb.HelloReply{Message: "hello " + in.GetName()}, nil
}

func (g *Greeter) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received params: %v", in.GetName())
	return &pb.HelloReply{Message: "hello again" + in.GetName()}, nil
}
