package grpclient

import (
	"ad-service-rpc/pb"
	"context"
	"log"
	"time"
)

func (c *Client) PersonList(request *pb.PersonListRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Person.PersonList(ctx, request)
	if err != nil {
		log.Fatalf("rpc could not conn: %v", err)
	}
	log.Printf("return: %v ", r.GetCount())
	log.Println(r.GetWomen())
}
