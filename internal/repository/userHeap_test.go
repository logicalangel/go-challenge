package repository

import (
	"arman-task/internal/models"
	"fmt"
	"testing"
	"time"
)

func TestUserHeap(t *testing.T) {
	userHeap := NewUserTtlHeap()

	for i := 0; i <= 10; i++ {
		userHeap.Push(&models.User{
			UserName:  fmt.Sprintf("user-%d", i),
			Segment:   "test-segment",
			CreatedAt: time.Now().Add(time.Hour * time.Duration(i)),
		})
	}

	if time.Now().Unix()-(userHeap).Head().CreatedAt.Unix() > int64(time.Hour) {
		t.Errorf("head of heap is not minimum")
	}
}
