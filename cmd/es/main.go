package main

import (
	"arman-task/internal/controller"
	"arman-task/internal/repository"
	"arman-task/internal/services/grpc"
	"arman-task/internal/usecase"
	"fmt"
)

func main() {
	segmentMap := repository.NewSegmentMap()
	userHeap := repository.NewUserTtlHeap()

	go grpc.NewServer(&userHeap, &segmentMap)
	go usecase.RunGarbageCollector(segmentMap, userHeap)

	fmt.Println("---- get segment estimate ----")
	for {
		controller.GetSegmentEstimate(segmentMap)
	}
}
