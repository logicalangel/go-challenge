package grpc

import (
	"arman-task/internal/usecase"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type EsServer struct {
	UnimplementedEsServiceServer
	segmentUsecase usecase.SegmentUseCase
}

func (es *EsServer) SendUserSegment(ctx context.Context, input *SendUserSegmentRequest) (*SendUserSegmentResponse, error) {
	err := es.segmentUsecase.SaveUserSegmention(input.Username, input.Segment)
	if err != nil {
		log.Fatal(err)
		return &SendUserSegmentResponse{}, err
	}

	return &SendUserSegmentResponse{
		Result: "ok",
	}, nil
}

func NewServer(segmentUsecase usecase.SegmentUseCase) {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()
	RegisterEsServiceServer(s, &EsServer{
		segmentUsecase: segmentUsecase,
	})
	if err := s.Serve(lis); err != nil {
		panic(err.Error())
	}
}
