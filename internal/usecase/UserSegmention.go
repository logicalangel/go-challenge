package usecase

import (
	"arman-task/internal/models"
	"arman-task/internal/repository"
	"time"
)

// these functions do everything about users and its segments and use services
type SegmentUseCase interface {
	SaveUserSegmention(username, segment string) error
	Estimate(segment string) uint64
}
type Segments struct {
	usersHeap  repository.UserHeapRepository
	segmentMap repository.SegmentMapRepository
}

func NewSegmentUsecase(usersHeap repository.UserHeapRepository, segmentMap repository.SegmentMapRepository) SegmentUseCase {
	return Segments{
		usersHeap:  usersHeap,
		segmentMap: segmentMap,
	}
}

func (s Segments) SaveUserSegmention(username, segment string) error {
	s.usersHeap.Push(&models.User{
		UserName:  username,
		Segment:   segment,
		CreatedAt: time.Now(),
	})

	s.segmentMap.IncreaseSegmentEstimate(segment)

	return nil
}

func (s Segments) Estimate(segment string) uint64 {
	return s.segmentMap.GetSegmentEstimate(segment)
}
