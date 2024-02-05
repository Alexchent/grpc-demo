package services

import (
	pb "ad-service-rpc/pb"
	"context"
	"log"
)

type Person struct {
	pb.UnimplementedPersonServer
}

func (svc *Person) PersonList(ctx context.Context, in *pb.PersonListRequest) (*pb.PersonListResponse, error) {
	log.Printf("Received params: %v, %s", in.GetAge(), in.GetName())

	girl1 := pb.Woman{
		Id:    2,
		Name:  "小乔",
		Age:   16,
		Level: pb.Woman_STUDENT,
	}

	girl2 := pb.Woman{
		Id:    1,
		Name:  "大乔",
		Age:   18,
		Level: pb.Woman_STUDENT,
	}
	girls := []*pb.Woman{&girl1, &girl2}
	return &pb.PersonListResponse{Count: 2, Women: girls}, nil
}
