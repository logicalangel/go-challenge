package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type EsServer struct {
	UnimplementedEsServiceServer
}

func (es *EsServer) SendUserSegment(ctx context.Context, input *SendUserSegmentRequest) (*SendUserSegmentResponse, error) {
	// TODO: call usecase.addUser
	return nil, nil
}

func NewServer() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()
	RegisterEsServiceServer(s, &EsServer{})
	if err := s.Serve(lis); err != nil {
		panic(err.Error())
	}
}
