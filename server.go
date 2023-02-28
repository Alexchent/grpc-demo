package main

import (
	pb "ad-service-rpc/pb"
	"ad-service-rpc/services"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The RpcSvc port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	//pb.RegisterEchoServer(s, &services.EchoSvc{})
	pb.RegisterGreeterServer(s, &services.Greeter{})
	pb.RegisterPersonServer(s, &services.Person{})

	log.Printf("GreeterSvc listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
