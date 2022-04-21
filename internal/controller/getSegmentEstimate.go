package controller

import (
	"arman-task/internal/usecase"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func GetSegmentEstimate(segmentUsecase usecase.SegmentUseCase) {
	promptSegment := promptui.Prompt{
		Label: "Enter Segment",
	}

	segment, err := promptSegment.Run()
	if err != nil {
		log.Fatal("failed to input segment: ", err.Error())
	}

	segmentEstimate := segmentUsecase.Estimate(segment)

	fmt.Printf("Segment estimate: %d\n------------------------------\n", segmentEstimate)
}
