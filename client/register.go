package grpclient

import "ad-service-rpc/pb"

type Client struct {
	Greeter pb.GreeterClient
	Person  pb.PersonClient
}
