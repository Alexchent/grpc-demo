package client

import (
	"ad-service-rpc/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var client pb.PersonClient

func init() {
	// Set up a connection to the GreeterSvc.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	// 初始化客户端
	client = pb.NewPersonClient(conn)
}

func PersonList(request *pb.PersonListRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.PersonList(ctx, request)
	if err != nil {
		log.Fatalf("rpc could not conn: %v", err)
	}
	log.Printf("return: %v ", r.GetCount())
	log.Println(r.GetWomen())
}
