package usecase

import (
	"arman-task/internal/repository"
	"time"
)

// this function run as a goroutine and check head of heap every minute
// if head of heap was expired pop and delete users until no expied user found

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
