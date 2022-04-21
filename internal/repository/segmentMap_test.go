package repository

import "testing"

func TestSegmentMap(t *testing.T) {
	sm := NewSegmentMap()

	for i := 0; i <= 10; i++ {
		sm.IncreaseSegmentEstimate("test-segment")
	}
	sm.DecreaseSegmentEstimate("test-segment")
	if sm.GetSegmentEstimate("test-segment") != 9 {
		t.Errorf("got %d, want: 9", sm.GetSegmentEstimate("test-segment"))
	}
}
