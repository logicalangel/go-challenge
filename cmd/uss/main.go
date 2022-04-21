package main

import (
	"arman-task/internal/controller"
	"arman-task/internal/repository"
	"arman-task/internal/services/grpc"
	"arman-task/internal/usecase"
	"fmt"
)

func main() {
	// create grpc client
	es := grpc.NewClient()

	// create repositories and usecase
	segmentMap := repository.NewSegmentMap()
	userHeap := repository.NewUserTtlHeap()
	segmentUsecase := usecase.NewSegmentUsecase(userHeap, segmentMap)

	// start user input
	fmt.Println("---- Add new user ----")
	for {
		controller.AddUserAndSegmention(segmentUsecase, es)
	}
}
