package grpc

import (
	"arman-task/internal/repository"
	"arman-task/internal/usecase"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type EsServer struct {
	UnimplementedEsServiceServer
	usersHeap  *repository.UserHeapRepository
	segmentMap *repository.SegmentMapRepository
}

func (es *EsServer) SendUserSegment(ctx context.Context, input *SendUserSegmentRequest) (*SendUserSegmentResponse, error) {
	err := usecase.SaveUserSegmention(*es.usersHeap, *es.segmentMap, input.Username, input.Segment)
	if err != nil {
		log.Fatal(err)
		return &SendUserSegmentResponse{}, err
	}

	return &SendUserSegmentResponse{
		Result: "ok",
	}, nil
}

func NewServer(usersHeap *repository.UserHeapRepository, segmentMap *repository.SegmentMapRepository) {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()
	RegisterEsServiceServer(s, &EsServer{
		usersHeap:  usersHeap,
		segmentMap: segmentMap,
	})
	if err := s.Serve(lis); err != nil {
		panic(err.Error())
	}
}
