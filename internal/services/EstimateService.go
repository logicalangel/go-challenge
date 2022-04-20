package services

type EstimateService interface {
	SendUserSegment(username, segment string) error
}
