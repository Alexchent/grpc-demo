package main

import (
	grpclient "ad-service-rpc/client"
	"ad-service-rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	// 初始化客户端
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
