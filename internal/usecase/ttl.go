package usecase

import (
	"arman-task/internal/repository"
	"time"
)

func RunGarbageCollector(sm repository.SegmentMapRepository, heap repository.UserHeapRepository) {
	for {
		heapHead := heap.Head()
		if heapHead != nil {
			for time.Now().Unix()-heapHead.CreatedAt.Unix() > int64(time.Hour*24*7) {
				user := heap.Pop()
				sm.DecreaseSegmentEstimate(user.Segment)
			}
		}
		time.Sleep(time.Minute)
	}
}
