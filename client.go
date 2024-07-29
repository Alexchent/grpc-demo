package main

import (
	grpclient "ad-service-rpc/client"
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
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "tony", "Name to greet")
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

	// 注册多个grpc客户端到client
	client := &grpclient.Client{}
	client.Greeter = pb.NewGreeterClient(conn)
	client.Person = pb.NewPersonClient(conn)

	client.PersonList(&pb.PersonListRequest{
		Name: "刀白凤",
		Age:  22,
	})

	client.SayHello(&pb.HelloRequest{Name: "小王"})
	client.SayHelloAgain(&pb.HelloRequest{Name: "小王"})
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
