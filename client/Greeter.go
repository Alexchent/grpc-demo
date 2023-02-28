package client

import (
	"ad-service-rpc/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var Greeter pb.GreeterClient

func init() {
	// Set up a connection to the GreeterSvc.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	// 初始化客户端
	Greeter = pb.NewGreeterClient(conn)
}

func SayHello(request *pb.HelloRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := Greeter.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("rpc could not conn: %v", err)
	}
	log.Printf("return: %v ", r.GetMessage())
	log.Println(r.GetMessage())
}

func SayHelloAgain(request *pb.HelloRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := Greeter.SayHelloAgain(ctx, request)
	if err != nil {
		log.Fatalf("rpc could not conn: %v", err)
	}
	log.Printf("return: %v ", r.GetMessage())
	log.Println(r.GetMessage())
}
