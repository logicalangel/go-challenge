package grpc

import (
	"arman-task/internal/services"
	"context"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	grpc EsServiceClient
}

func (gc GrpcClient) SendUserSegment(username, segment string) error {
	_, err := gc.grpc.SendUserSegment(context.TODO(), &SendUserSegmentRequest{
		Username: username,
		Segment:  segment,
	})
	return err
}

func NewClient() services.EstimateService {
	es := GrpcClient{}
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err.Error())
	}

	es.grpc = NewEsServiceClient(conn)
	return es
}
