package repository

import (
	"arman-task/internal/models"
	"container/heap"
)

type UserTtlHeap []*models.User

func (h UserTtlHeap) Len() int {
	return len(h)
}
func (h UserTtlHeap) Less(i, j int) bool {
	return h[i].CreatedAt.Unix() < h[j].CreatedAt.Unix()
}
func (h UserTtlHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *UserTtlHeap) Push(x interface{}) {
	*h = append(*h, x.(*models.User))
}

func (h *UserTtlHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *UserTtlHeap) Peek() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}

type UserHeapRepository interface {
	Len() int
	Head() *models.User
	Push(user *models.User)
	Pop() *models.User
}

type UserHeap struct {
	heap UserTtlHeap
}

func (uh UserHeap) Len() int {
	return uh.heap.Len()
}

func (uh UserHeap) Head() *models.User {
	if len(uh.heap) > 0 {
		return uh.heap[0]
	} else {
		return nil
	}
}

func (uh UserHeap) Push(user *models.User) {
	heap.Push(&uh.heap, user)
}

func (uh UserHeap) Pop() *models.User {
	return uh.heap.Pop().(*models.User)
}

func NewUserTtlHeap() UserHeapRepository {
	var ttlHeap *UserTtlHeap = &UserTtlHeap{}
	heap.Init(ttlHeap)

	return UserHeap{
		heap: *ttlHeap,
	}
}
