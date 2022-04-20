package controller

import (
	"arman-task/internal/repository"
	"arman-task/internal/usecase"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func GetSegmentEstimate(sm repository.SegmentMapRepository) {
	promptSegment := promptui.Prompt{
		Label: "Enter Segment",
	}

	segment, err := promptSegment.Run()
	if err != nil {
		log.Fatal("failed to input segment: ", err.Error())
	}

	segmentEstimate := usecase.Estimate(sm, segment)

	fmt.Printf("Segment estimate: %d\n------------------------------\n", segmentEstimate)
}
