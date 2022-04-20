package repository

type SegmentMapRepository interface {
	IncreaseSegmentEstimate(segment string)
	DecreaseSegmentEstimate(segment string)
	GetSegmentEstimate(segment string) uint64
}

type SegmentMap struct {
	segmentMap map[string]uint64
}

func NewSegmentMap() SegmentMapRepository {
	return &SegmentMap{segmentMap: map[string]uint64{}}
}

func (sm SegmentMap) IncreaseSegmentEstimate(segment string) {
	sm.segmentMap[segment]++
}

func (sm SegmentMap) DecreaseSegmentEstimate(segment string) {
	sm.segmentMap[segment]--
}

func (sm SegmentMap) GetSegmentEstimate(segment string) uint64 {
	return sm.segmentMap[segment]
}
