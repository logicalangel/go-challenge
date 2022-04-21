package main

import (
	"arman-task/internal/controller"
	"arman-task/internal/repository"
	"arman-task/internal/services/grpc"
	"arman-task/internal/usecase"
	"fmt"
)

func main() {
	// create repositories
	segmentMap := repository.NewSegmentMap()
	userHeap := repository.NewUserTtlHeap()

	// run services
	go grpc.NewServer(&userHeap, &segmentMap)
	go usecase.RunGarbageCollector(segmentMap, userHeap)

	// start user input
	fmt.Println("---- get segment estimate ----")
	for {
		controller.GetSegmentEstimate(segmentMap)
	}
}
