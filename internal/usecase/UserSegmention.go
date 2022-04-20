package usecase

import (
	"arman-task/internal/models"
	"arman-task/internal/repository"
	"arman-task/internal/services"
	"time"
)

func SendUserSegmention(es services.EstimateService, username, segment string) error {
	err := es.SendUserSegment(username, segment)
	if err != nil {
		return err
	}

	return nil
}

func SaveUserSegmention(usersHeap repository.UserHeapRepository, segmentMap repository.SegmentMapRepository, username, segment string) error {
	usersHeap.Push(&models.User{
		UserName:  username,
		Segment:   segment,
		CreatedAt: time.Now(),
	})

	segmentMap.IncreaseSegmentEstimate(segment)

	return nil
}

func Estimate(sm repository.SegmentMapRepository, segment string) uint64 {
	return sm.GetSegmentEstimate(segment)
}
