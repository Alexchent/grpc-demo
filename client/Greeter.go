package grpclient

import (
	"ad-service-rpc/pb"
	"context"
	"log"
	"time"
)

func (c *Client) SayHello(request *pb.HelloRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Greeter.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("rpc could not conn: %v", err)
	}
	log.Printf("return: %v ", r.GetMessage())
	log.Println(r.GetMessage())
}

func (c *Client) SayHelloAgain(request *pb.HelloRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Greeter.SayHelloAgain(ctx, request)
	if err != nil {
		log.Fatalf("rpc could not conn: %v", err)
	}
	log.Printf("return: %v ", r.GetMessage())
	log.Println(r.GetMessage())
}
