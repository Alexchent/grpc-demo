package main

import (
	"ad-service-rpc/pb"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// go run client.go -name=coo
var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect to")
	name    = flag.String("name", "tony", "Name to greet")
	message = flag.String("message", "我的女人", "message to echo")
)

func main() {
	flag.Parse()
	// Set up a connection to the GreeterSvc.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 初始化客户端
	Hello(pb.NewGreeterClient(conn), &pb.HelloRequest{Name: *name})
}

func Hello(c pb.GreeterClient, request *pb.HelloRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("GreeterSvc return: " + r.GetMessage())
}
