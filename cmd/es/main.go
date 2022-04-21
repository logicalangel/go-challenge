package main

import (
	"arman-task/internal/controller"
	"arman-task/internal/repository"
	"arman-task/internal/services/grpc"
	"arman-task/internal/usecase"
	"fmt"
)

func main() {
	// create repositories and usecase
	segmentMap := repository.NewSegmentMap()
	userHeap := repository.NewUserTtlHeap()
	segmentUsecase := usecase.NewSegmentUsecase(userHeap, segmentMap)

	// run services
	go grpc.NewServer(segmentUsecase)
	go usecase.RunGarbageCollector(segmentMap, userHeap)

	// start user input
	fmt.Println("---- get segment estimate ----")
	for {
		controller.GetSegmentEstimate(segmentUsecase)
	}
}
