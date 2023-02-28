package main

import (
	"ad-service-rpc/client"
	"ad-service-rpc/pb"
)

func main() {
	client.PersonList(&pb.PersonListRequest{
		Name: "刀白凤",
		Age:  22,
	})

	client.SayHello(&pb.HelloRequest{Name: "小王"})
	client.SayHelloAgain(&pb.HelloRequest{Name: "小王"})
}
